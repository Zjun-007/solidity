package repository

import (
    "gorm.io/gorm"
    "geth/internal/model"
)

type AuctionRepository struct {
    db *gorm.DB
}

func NewAuctionRepository(db *gorm.DB) *AuctionRepository {
    return &AuctionRepository{db: db}
}

func (r *AuctionRepository) CreateAuction(auction *model.Auction) error {
    return r.db.Create(auction).Error
}

func (r *AuctionRepository) GetAuctionByID(auctionID uint64) (*model.Auction, error) {
    var auction model.Auction
    err := r.db.Where("auction_id = ?", auctionID).First(&auction).Error
    return &auction, err
}

func (r *AuctionRepository) ListActiveAuctions() ([]model.Auction, error) {
    var auctions []model.Auction
    err := r.db.Where("ended = ? AND end_time > NOW()", false).Find(&auctions).Error
    return auctions, err
}

func (r *AuctionRepository) ListAllAuctions() ([]model.Auction, error) {
    var auctions []model.Auction
    err := r.db.Order("created_at DESC").Find(&auctions).Error
    return auctions, err
}

func (r *AuctionRepository) CreateBid(bid *model.Bid) error {
    return r.db.Create(bid).Error
}

func (r *AuctionRepository) UpdateAuction(auctionID uint64, updates map[string]interface{}) error {
    return r.db.Model(&model.Auction{}).Where("auction_id = ?", auctionID).Updates(updates).Error
}

func (r *AuctionRepository) GetBidsByAuction(auctionID uint64) ([]model.Bid, error) {
    var bids []model.Bid
    err := r.db.Where("auction_id = ?", auctionID).Order("created_at DESC").Find(&bids).Error
    return bids, err
}