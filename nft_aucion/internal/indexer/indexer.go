package indexer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"

	"geth/internal/blockchain"
	"geth/internal/config"
	"geth/internal/model"
	"geth/internal/repository"
)

// 与 forge inspect 一致的事件 topic0
var (
	topicAuctionCreated  = common.HexToHash("0xcaf0ae751fb2b122e8718bf7c0d4b7584d1418a853a4d0cdaba45418d3da138b")
	topicBidAuction      = common.HexToHash("0x014de1dd43627559120d21162d1154078a340acf94db1e6c44d3f1933ff03df9")
	topicAuctionEnded    = common.HexToHash("0xbb6764412c29916bdf4a5c6fe6b1c079de35682160b2289928ce003ab459a749")
	topicAuctionCanceled = common.HexToHash("0x10ac9f0bb365b5d22d7bec500408692f23fdf83eadfec71615ef88b4c1134f0e")
)

// Indexer watches the NFT auction proxy contract for relevant events and writes
// the auction state into the local database. It also exposes a lightweight
// status API so the HTTP layer can report sync progress.
type Indexer struct {
	cfg     *config.Config
	eth     *ethclient.Client
	addr    common.Address
	filter  *blockchain.NFTAuctionUUPSFilterer
	caller  *blockchain.NFTAuctionUUPSCaller
	baseDB  *gorm.DB
	store   *repository.UUPSStore

	mu         sync.RWMutex
	lastSynced uint64
	chainHead  uint64
	lastErr    string
}

// New creates a new Indexer instance with an Ethereum client, event filterer, and
// contract caller bound to the configured proxy address.
func New(cfg *config.Config, db *gorm.DB, cli *ethclient.Client) (*Indexer, error) {
	addr := common.HexToAddress(cfg.NFTAuctionProxy)
	f, err := blockchain.NewNFTAuctionUUPSFilterer(addr, cli)
	if err != nil {
		return nil, err
	}
	c, err := blockchain.NewNFTAuctionUUPSCaller(addr, cli)
	if err != nil {
		return nil, err
	}
	return &Indexer{
		cfg:    cfg,
		eth:    cli,
		addr:   addr,
		filter: f,
		caller: c,
		baseDB: db,
		store:  repository.NewUUPSStore(db),
	}, nil
}

func (ix *Indexer) Status() (lastSynced, head uint64, lastErr string) {
	ix.mu.RLock()
	defer ix.mu.RUnlock()
	return ix.lastSynced, ix.chainHead, ix.lastErr
}

func (ix *Indexer) setStatus(synced, head uint64, errStr string) {
	ix.mu.Lock()
	defer ix.mu.Unlock()
	ix.lastSynced = synced
	ix.chainHead = head
	ix.lastErr = errStr
}

// Run starts the indexer loop. It polls the chain at a configured interval and
// continues until the provided context is canceled.
func (ix *Indexer) Run(ctx context.Context) {
	log.Printf("[indexer] starting with poll interval %v, chunk size %d blocks", ix.cfg.IndexPollInterval, ix.cfg.IndexChunkBlocks)
	t := time.NewTicker(ix.cfg.IndexPollInterval)
	defer t.Stop()
	for {
		if err := ix.poll(ctx); err != nil {
			log.Printf("[indexer] poll error: %v", err)
			ix.mu.Lock()
			ix.lastErr = err.Error()
			ix.mu.Unlock()
		}
		select {
		case <-ctx.Done():
			log.Printf("[indexer] stopping due to context cancellation")
			return
		case <-t.C:
		}
	}
}

// poll reads the newest chain head, queries logs for the next event range, then
// processes each log inside a DB transaction and advances the cursor.
func (ix *Indexer) poll(ctx context.Context) error {
	head, err := ix.eth.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain head: %w", err)
	}
	ix.mu.Lock()
	ix.chainHead = head
	ix.mu.Unlock()

	from, ok, err := ix.store.GetCursor(model.CursorKeyNFTAuction)
	if err != nil {
		return fmt.Errorf("failed to get cursor: %w", err)
	}
	if !ok {
		from = ix.cfg.AuctionDeployFromBlock
		if from > 0 {
			from--
		}
		log.Printf("[indexer] initializing cursor to block %d", from)
	}
	start := from + 1
	if start > head {
		ix.setStatus(from, head, "")
		log.Printf("[indexer] no new blocks to process (head: %d, synced: %d)", head, from)
		return nil
	}

	to := start + ix.cfg.IndexChunkBlocks - 1
	if to > head {
		to = head
	}

	log.Printf("[indexer] processing blocks %d to %d", start, to)
	topics := [][]common.Hash{{topicAuctionCreated, topicBidAuction, topicAuctionEnded, topicAuctionCanceled}}
	logs, err := ix.eth.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(to),
		Addresses: []common.Address{ix.addr},
		Topics:    topics,
	})
	if err != nil {
		return fmt.Errorf("failed to filter logs: %w", err)
	}

	if len(logs) > 0 {
		log.Printf("[indexer] found %d events to process", len(logs))
	}

	err = ix.baseDB.Transaction(func(tx *gorm.DB) error {
		st := repository.NewUUPSStore(tx)
		for i := range logs {
			if e := ix.processLog(ctx, st, &logs[i]); e != nil {
				return fmt.Errorf("failed to process log at index %d: %w", i, e)
			}
		}
		return st.SetCursor(model.CursorKeyNFTAuction, to)
	})
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}
	ix.setStatus(to, head, "")
	log.Printf("[indexer] synced to block %d", to)
	return nil
}

// processLog decodes a contract event and updates the database accordingly.
func (ix *Indexer) processLog(ctx context.Context, st *repository.UUPSStore, lg *types.Log) error {
	if len(lg.Topics) == 0 {
		return nil
	}
	switch lg.Topics[0] {
	case topicAuctionCreated:
		ev, err := ix.filter.ParseAuctionCreated(*lg)
		if err != nil {
			return err
		}
		a := &model.UUPSAuction{
			AuctionID:     ev.AuctionId.Uint64(),
			Seller:        addrLower(ev.Seller),
			NFTContract:   addrLower(ev.NftContract),
			TokenID:       ev.TokenId.String(),
			StartPrice:    ev.StartPrice.String(),
			StartTime:     ev.StartTime.Uint64(),
			EndTime:       ev.EndTime.Uint64(),
			HighestBidder: "",
			HighestBid:    "0",
			CreatedBlock:  lg.BlockNumber,
		}
		return st.UpsertAuctionCreated(a)

	case topicBidAuction:
		ev, err := ix.filter.ParseBidAuction(*lg)
		if err != nil {
			return err
		}
		b := &model.UUPSBid{
			AuctionID:   ev.AuctionId.Uint64(),
			Bidder:      addrLower(ev.Bidder),
			Amount:      ev.Amount.String(),
			BlockNumber: lg.BlockNumber,
			TxHash:      strings.ToLower(lg.TxHash.Hex()),
			LogIndex:    uint(lg.Index),
		}
		inserted, err := st.InsertBid(b)
		if err != nil {
			return err
		}
		if !inserted {
			return nil
		}
		return st.UpdateAuctionBid(b.AuctionID, b.Bidder, b.Amount, lg.BlockNumber)

	case topicAuctionEnded:
		ev, err := ix.filter.ParseAuctionEnded(*lg)
		if err != nil {
			return err
		}
		return st.MarkAuctionEnded(ev.AuctionId.Uint64(), lg.BlockNumber)

	case topicAuctionCanceled:
		ev, err := ix.filter.ParseAuctionCancelled(*lg)
		if err != nil {
			return err
		}
		return st.MarkAuctionCancelled(ev.AuctionId.Uint64(), lg.BlockNumber)

	default:
		return nil
	}
}

// reconcileChain queries the contract directly for auctions that may be stale or
// out of sync, ensuring that open auctions are corrected from on-chain state.
func (ix *Indexer) reconcileChain(ctx context.Context, head uint64) error {
	now := time.Now().Unix()
	rows, err := ix.store.ListStaleOpenAuctions(uint64(now), 80)
	if err != nil {
		return err
	}
	opts := &bind.CallOpts{Context: ctx}
	for i := range rows {
		id := new(big.Int).SetUint64(rows[i].AuctionID)
		data, err := ix.caller.AuctionData(opts, id)
		if err != nil {
			log.Printf("[indexer] reconcile auctionData(%d): %v", rows[i].AuctionID, err)
			continue
		}
		if data.Seller == (common.Address{}) {
			continue
		}
		hb := "0"
		hbder := ""
		if data.HighestBid != nil {
			hb = data.HighestBid.String()
		}
		if data.HighestBidder != (common.Address{}) {
			hbder = addrLower(data.HighestBidder)
		}
		a := &model.UUPSAuction{
			AuctionID:     rows[i].AuctionID,
			Seller:        addrLower(data.Seller),
			NFTContract:   addrLower(data.NftContract),
			TokenID:       data.TokenId.String(),
			StartPrice:    data.StartPrice.String(),
			StartTime:     data.StartTime.Uint64(),
			EndTime:       data.EndTime.Uint64(),
			HighestBidder: hbder,
			HighestBid:    hb,
			Settled:       data.Settled,
			Cancelled:     data.Cancelled,
			CreatedBlock:  rows[i].CreatedBlock,
			UpdatedBlock:  head,
		}
		if err := ix.store.SyncAuctionFromChain(a); err != nil {
			log.Printf("[indexer] reconcile save %d: %v", rows[i].AuctionID, err)
		}
	}
	return nil
}

func addrLower(a common.Address) string {
	return strings.ToLower(a.Hex())
}
