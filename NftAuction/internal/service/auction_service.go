package service

import (
    "context"
    "errors"
    "math/big"

    "geth/internal/blockchain"
)

type AuctionService struct {
    bcClient *blockchain.Client
}

func NewAuctionService(bc *blockchain.Client) *AuctionService {
    return &AuctionService{bcClient: bc}
}

type BidRequest struct {
    Amount string `json:"amount" binding:"required"` // 单位 Wei
}

// DepositNFT 卖家存入 NFT
func (s *AuctionService) DepositNFT() (string, error) {
    tx, err := s.bcClient.DepositNFT()
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

// PlaceBid 出价
func (s *AuctionService) PlaceBid(amount string) (string, error) {
    bidAmount, ok := new(big.Int).SetString(amount, 10)
    if !ok {
        return "", errors.New("invalid bid amount")
    }
    tx, err := s.bcClient.Bid(bidAmount)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

// EndAuction 结束拍卖
func (s *AuctionService) EndAuction() (string, error) {
    tx, err := s.bcClient.EndAuction()
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

// Withdraw 提取被超越的出价
func (s *AuctionService) Withdraw() (string, error) {
    tx, err := s.bcClient.Withdraw()
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

// CancelAuction 取消拍卖
func (s *AuctionService) CancelAuction() (string, error) {
    tx, err := s.bcClient.CancelAuction()
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

// GetAuctionInfo 获取当前拍卖信息
func (s *AuctionService) GetAuctionInfo(ctx context.Context) (*blockchain.AuctionInfo, error) {
    return s.bcClient.GetAuctionInfo(ctx)
}