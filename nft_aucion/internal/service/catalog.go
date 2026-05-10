package service

import (
	"context"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"geth/internal/model"
	"geth/internal/nft"
	"geth/internal/repository"
)

type Catalog struct {
	store *repository.UUPSStore
	eth   *ethclient.Client
}

func NewCatalog(store *repository.UUPSStore, eth *ethclient.Client) *Catalog {
	return &Catalog{store: store, eth: eth}
}

type AuctionListQuery struct {
	Status      string `form:"status"`
	Seller      string `form:"seller"`
	NFTContract string `form:"nft_contract"`
	Sort        string `form:"sort"`
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
}

func (c *Catalog) ListAuctions(q AuctionListQuery) ([]model.UUPSAuction, int64, error) {
	return c.store.ListAuctions(repository.ListAuctionParams{
		Status:      q.Status,
		Seller:      strings.TrimSpace(q.Seller),
		NFTContract: strings.TrimSpace(q.NFTContract),
		Sort:        q.Sort,
		Page:        q.Page,
		PageSize:    q.PageSize,
	})
}

func (c *Catalog) ListBids(auctionID uint64, page, pageSize int) ([]model.UUPSBid, int64, error) {
	return c.store.ListBidsByAuction(auctionID, page, pageSize)
}

func (c *Catalog) Stats() (auctions, bids int64, err error) {
	return c.store.Stats()
}

func (c *Catalog) WalletNFTs(ctx context.Context, wallet string, contractsCSV string) ([]nft.ContractTokens, error) {
	owner := common.HexToAddress(wallet)
	var parts []string
	for _, p := range strings.Split(contractsCSV, ",") {
		if t := strings.TrimSpace(p); t != "" {
			parts = append(parts, t)
		}
	}
	if len(parts) == 0 {
		return nil, errors.New("contracts query param required (comma-separated ERC721 addresses)")
	}
	return nft.ListTokensByContracts(ctx, c.eth, owner, parts)
}
