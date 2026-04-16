// internal/blockchain/event_listener.go
package blockchain

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "gorm.io/gorm"
    "geth/internal/model"  // 请替换为你的实际模型路径
)

type EventListener struct {
    client       *ethclient.Client
    contractAddr common.Address
    db           *gorm.DB
    contract     *Blockchain          // 改为 Blockchain
}

func NewEventListener(wsURL, contractAddr string, db *gorm.DB) (*EventListener, error) {
    client, err := ethclient.Dial(wsURL)
    if err != nil {
        return nil, fmt.Errorf("dial ws: %w", err)
    }
    contractAddress := common.HexToAddress(contractAddr)
    contract, err := NewBlockchain(contractAddress, client)   // 改为 NewBlockchain
    if err != nil {
        return nil, fmt.Errorf("bind contract: %w", err)
    }
    return &EventListener{
        client:       client,
        contractAddr: contractAddress,
        db:           db,
        contract:     contract,
    }, nil
}

func (el *EventListener) Start() {
    query := ethereum.FilterQuery{
        Addresses: []common.Address{el.contractAddr},
    }
    logs := make(chan types.Log)
    sub, err := el.client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal("Subscription error:", err)
    }
    for {
        select {
        case err := <-sub.Err():
            log.Println("Subscription error, reconnecting:", err)
            time.Sleep(5 * time.Second)
            continue
        case vLog := <-logs:
            el.processLog(vLog)
        }
    }
}

func (el *EventListener) processLog(vLog types.Log) {
    // 防止重复处理
    var syncLog model.SyncLog
    result := el.db.Where("tx_hash = ? AND event_type = ?", vLog.TxHash.Hex(), vLog.Topics[0].Hex()).First(&syncLog)
    if result.Error == nil {
        log.Printf("Event already processed: tx=%s", vLog.TxHash.Hex())
        return
    }

    contractAbi, err := BlockchainMetaData.GetAbi()   // 改为 BlockchainMetaData
    if err != nil {
        log.Printf("Failed to get ABI: %v", err)
        return
    }

    // 匹配事件签名
    for name, event := range contractAbi.Events {
        if vLog.Topics[0].Hex() == event.ID.Hex() {
            switch name {
            case "AuctionCreated":
                var ev BlockchainAuctionCreated
                if err := contractAbi.UnpackIntoInterface(&ev, name, vLog.Data); err != nil {
                    log.Printf("Unpack error: %v", err)
                    return
                }
                // 从 topics 中提取 indexed 参数
                if len(vLog.Topics) > 1 {
                    ev.Seller = common.BytesToAddress(vLog.Topics[1].Bytes())
                }
                el.handleAuctionCreated(ev, vLog)
            case "Bid":
                var ev BlockchainBid
                if err := contractAbi.UnpackIntoInterface(&ev, name, vLog.Data); err != nil {
                    log.Printf("Unpack error: %v", err)
                    return
                }
                if len(vLog.Topics) > 1 {
                    ev.Bidder = common.BytesToAddress(vLog.Topics[1].Bytes())
                }
                el.handleBid(ev, vLog)
            case "AuctionEnded":
                var ev BlockchainAuctionEnded
                if err := contractAbi.UnpackIntoInterface(&ev, name, vLog.Data); err != nil {
                    log.Printf("Unpack error: %v", err)
                    return
                }
                el.handleAuctionEnded(ev, vLog)
            case "AuctionCanceled":
                el.handleAuctionCanceled(vLog)
            case "NFTDeposited":
                el.handleNFTDeposited(vLog)
            case "Withdraw":
                var ev BlockchainWithdraw
                if err := contractAbi.UnpackIntoInterface(&ev, name, vLog.Data); err != nil {
                    log.Printf("Unpack error: %v", err)
                    return
                }
                if len(vLog.Topics) > 1 {
                    ev.Bidder = common.BytesToAddress(vLog.Topics[1].Bytes())
                }
                el.handleWithdraw(ev, vLog)
            default:
                log.Printf("Unknown event: %s", name)
            }
            // 记录已处理
            el.db.Create(&model.SyncLog{
                EventType:   vLog.Topics[0].Hex(),
                BlockNumber: vLog.BlockNumber,
                TxHash:      vLog.TxHash.Hex(),
                ProcessedAt: time.Now(),
            })
            return
        }
    }
    log.Printf("Unhandled topic: %s", vLog.Topics[0].Hex())
}

// 事件处理函数（示例）
func (el *EventListener) handleAuctionCreated(ev BlockchainAuctionCreated, vLog types.Log) {
    log.Printf("AuctionCreated: seller=%s, tokenId=%v, startPrice=%v, endTime=%v",
        ev.Seller.Hex(), ev.TokenId, ev.StartPrice, ev.EndTime)
    auction := &model.Auction{
        AuctionID:   0,  // 单次拍卖使用固定 ID
        NFTContract: el.contractAddr.Hex(),
        TokenID:     ev.TokenId.Uint64(),
        StartPrice:  ev.StartPrice.String(),
        EndTime:     time.Unix(ev.EndTime.Int64(), 0),
        Creator:     ev.Seller.Hex(),
        Ended:       false,
    }
    if err := el.db.Create(auction).Error; err != nil {
        log.Printf("Failed to save auction: %v", err)
    }
}

func (el *EventListener) handleBid(ev BlockchainBid, vLog types.Log) {
    log.Printf("Bid: bidder=%s, amount=%v", ev.Bidder.Hex(), ev.Amount)
    el.db.Model(&model.Auction{}).Where("auction_id = ?", 0).Updates(map[string]interface{}{
        "highest_bid":    ev.Amount.String(),
        "highest_bidder": ev.Bidder.Hex(),
    })
    bid := &model.Bid{
        AuctionID:   0,
        Bidder:      ev.Bidder.Hex(),
        Amount:      ev.Amount.String(),
        TxHash:      vLog.TxHash.Hex(),
        BlockNumber: vLog.BlockNumber,
    }
    el.db.Create(bid)
}

func (el *EventListener) handleAuctionEnded(ev BlockchainAuctionEnded, vLog types.Log) {
    log.Printf("AuctionEnded: winner=%s, amount=%v", ev.Winner.Hex(), ev.Amount)
    el.db.Model(&model.Auction{}).Where("auction_id = ?", 0).Update("ended", true)
}

func (el *EventListener) handleAuctionCanceled(vLog types.Log) {
    log.Println("AuctionCanceled")
    el.db.Model(&model.Auction{}).Where("auction_id = ?", 0).Update("ended", true) // 或者标记 canceled 字段
}

func (el *EventListener) handleNFTDeposited(vLog types.Log) {
    log.Println("NFTDeposited")
    // 可根据需要更新 nft_deposited 状态
}

func (el *EventListener) handleWithdraw(ev BlockchainWithdraw, vLog types.Log) {
    log.Printf("Withdraw: bidder=%s, amount=%v", ev.Bidder.Hex(), ev.Amount)
}