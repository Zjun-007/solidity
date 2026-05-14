package service

import (
    "errors"
    "math/big"
    "time"

    "github.com/ethereum/go-ethereum/common"
    "geth/internal/blockchain"
    "geth/internal/model"
    "geth/internal/repository"
)

type AuctionService struct {
    bcClient *blockchain.Client
    repo     *repository.AuctionRepository
}

func NewAuctionService(bc *blockchain.Client, repo *repository.AuctionRepository) *AuctionService {
    return &AuctionService{bcClient: bc, repo: repo}
}

type CreateAuctionRequest struct {
    NFTContract string `json:"nft_contract" binding:"required"`
    TokenID     uint64 `json:"token_id" binding:"required"`
    StartPrice  string `json:"start_price" binding:"required"` // 单位：Wei，字符串格式
    Duration    uint64 `json:"duration" binding:"required"`    // 秒
}

type BidRequest struct {
    Amount string `json:"amount" binding:"required"` // 单位：Wei
}

func (s *AuctionService) CreateAuction(req CreateAuctionRequest, creatorAddress string) (txHash string, err error) {
    nftContract := common.HexToAddress(req.NFTContract)
    tokenId := new(big.Int).SetUint64(req.TokenID)
    startPrice, ok := new(big.Int).SetString(req.StartPrice, 10)
    if !ok {
        return "", errors.New("invalid start price")
    }
    duration := new(big.Int).SetUint64(req.Duration)

    tx, err := s.bcClient.CreateAuction(nftContract, tokenId, startPrice, duration)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (s *AuctionService) PlaceBid(auctionID uint64, bidderAddress string, amount string) (txHash string, err error) {
    // 先检查拍卖是否活跃（从数据库快速检查）
    auction, err := s.repo.GetAuctionByID(auctionID)
    if err != nil {
        return "", errors.New("auction not found")
    }
    if auction.Ended {
        return "", errors.New("auction already ended")
    }
    if time.Now().After(auction.EndTime) {
        return "", errors.New("auction has ended")
    }

    bidAmount, ok := new(big.Int).SetString(amount, 10)
    if !ok {
        return "", errors.New("invalid bid amount")
    }

    // 可选：检查出价是否高于当前最高价（可略，合约会检查）
    tx, err := s.bcClient.Bid(new(big.Int).SetUint64(auctionID), bidAmount)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (s *AuctionService) EndAuction(auctionID uint64) (txHash string, err error) {
    auction, err := s.repo.GetAuctionByID(auctionID)
    if err != nil {
        return "", errors.New("auction not found")
    }
    if auction.Ended {
        return "", errors.New("auction already ended")
    }
    if time.Now().Before(auction.EndTime) {
        return "", errors.New("auction not ended yet")
    }

    tx, err := s.bcClient.EndAuction(new(big.Int).SetUint64(auctionID))
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (s *AuctionService) GetAuction(auctionID uint64) (*model.Auction, error) {
    return s.repo.GetAuctionByID(auctionID)
}

func (s *AuctionService) ListAuctions(activeOnly bool) ([]model.Auction, error) {
    if activeOnly {
        return s.repo.ListActiveAuctions()
    }
    return s.repo.ListAllAuctions()
}

func (s *AuctionService) GetBids(auctionID uint64) ([]model.Bid, error) {
    return s.repo.GetBidsByAuction(auctionID)
}