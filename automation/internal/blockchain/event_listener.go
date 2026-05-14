package blockchain

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "time"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "gorm.io/gorm"
    "geth/internal/model"
)

type EventListener struct {
    client           *ethclient.Client
    contractAddr     common.Address
    db               *gorm.DB
    contract         *FeedAuction
    eventSignatures  map[string]string
}

func NewEventListener(wsURL, contractAddr string, db *gorm.DB) (*EventListener, error) {
    client, err := ethclient.Dial(wsURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to WebSocket: %v", err)
    }

    contractAddress := common.HexToAddress(contractAddr)
    contract, err := NewFeedAuction(contractAddress, client)
    if err != nil {
        return nil, err
    }

    // 预计算事件签名
    eventSignatures := map[string]string{
        "AuctionCreated": "0x5d6b4e9e5f0f3d0d0c9a8e7f6e5d4c3b2a1908f7e6d5c4b3a291807f6e5d4c3b2a1",
        "NewBid":         "0x2b5d6c8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c",
        "AuctionEnded":   "0x7a6b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7",
    }

    return &EventListener{
        client:          client,
        contractAddr:    contractAddress,
        db:              db,
        contract:        contract,
        eventSignatures: eventSignatures,
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
    // 检查是否已处理
    var syncLog model.SyncLog
    result := el.db.Where("event_type = ? AND tx_hash = ?", vLog.Topics[0].Hex(), vLog.TxHash.Hex()).First(&syncLog)
    if result.Error == nil {
        log.Printf("Event already processed: tx=%s", vLog.TxHash.Hex())
        return
    }

    eventSig := vLog.Topics[0].Hex()

    switch eventSig {
    case el.eventSignatures["AuctionCreated"]:
        el.handleAuctionCreated(vLog)
    case el.eventSignatures["NewBid"]:
        el.handleNewBid(vLog)
    case el.eventSignatures["AuctionEnded"]:
        el.handleAuctionEnded(vLog)
    default:
        log.Printf("Unknown event signature: %s", eventSig)
    }

    // 记录处理日志
    syncLog = model.SyncLog{
        EventType:   eventSig,
        BlockNumber: vLog.BlockNumber,
        TxHash:      vLog.TxHash.Hex(),
        ProcessedAt: time.Now(),
    }
    el.db.Create(&syncLog)
}

func (el *EventListener) handleAuctionCreated(vLog types.Log) {
    if len(vLog.Topics) < 4 {
        log.Printf("Invalid AuctionCreated event: insufficient topics")
        return
    }

    auctionId := new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
    nftFeedAuction := common.BytesToAddress(vLog.Topics[2].Bytes()).Hex()
    creator := common.BytesToAddress(vLog.Topics[3].Bytes()).Hex()

    data := vLog.Data
    tokenId := new(big.Int).SetBytes(data[0:32]).Uint64()
    startPrice := new(big.Int).SetBytes(data[32:64])
    endTime := new(big.Int).SetBytes(data[64:96])

    auction := &model.Auction{
        AuctionID:   auctionId,
        NFTContract: nftFeedAuction,
        TokenID:     tokenId,
        StartPrice:  startPrice.String(),
        HighestBid:  "0",
        EndTime:     time.Unix(endTime.Int64(), 0),
        Ended:       false,
        Creator:     creator,
    }

    if err := el.db.Create(auction).Error; err != nil {
        log.Printf("Failed to save auction: %v", err)
    } else {
        log.Printf("Auction created: id=%d, nft=%s, token=%d", auctionId, nftFeedAuction, tokenId)
    }
}

func (el *EventListener) handleNewBid(vLog types.Log) {
    if len(vLog.Topics) < 3 {
        log.Printf("Invalid NewBid event: insufficient topics")
        return
    }

    auctionId := new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
    bidder := common.BytesToAddress(vLog.Topics[2].Bytes()).Hex()
    amount := new(big.Int).SetBytes(vLog.Data)

    bid := &model.Bid{
        AuctionID:   auctionId,
        Bidder:      bidder,
        Amount:      amount.String(),
        TxHash:      vLog.TxHash.Hex(),
        BlockNumber: vLog.BlockNumber,
    }

    if err := el.db.Create(bid).Error; err != nil {
        log.Printf("Failed to save bid: %v", err)
    }

    // 更新拍卖表的最高出价
    el.db.Model(&model.Auction{}).Where("auction_id = ?", auctionId).Updates(map[string]interface{}{
        "highest_bid":    amount.String(),
        "highest_bidder": bidder,
    })

    log.Printf("New bid: auction=%d, bidder=%s, amount=%s", auctionId, bidder, amount.String())
}

func (el *EventListener) handleAuctionEnded(vLog types.Log) {
    if len(vLog.Topics) < 2 {
        log.Printf("Invalid AuctionEnded event: insufficient topics")
        return
    }

    auctionId := new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
    winner := common.BytesToAddress(vLog.Data[0:32]).Hex()
    winningBid := new(big.Int).SetBytes(vLog.Data[32:64])

    el.db.Model(&model.Auction{}).Where("auction_id = ?", auctionId).Updates(map[string]interface{}{
        "ended": true,
    })

    log.Printf("Auction ended: id=%d, winner=%s, winningBid=%s", auctionId, winner, winningBid.String())
}