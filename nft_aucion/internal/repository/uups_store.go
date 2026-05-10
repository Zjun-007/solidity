package repository

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"geth/internal/model"
)

type UUPSStore struct {
	db *gorm.DB
}

func NewUUPSStore(db *gorm.DB) *UUPSStore {
	return &UUPSStore{db: db}
}

func (s *UUPSStore) GetCursor(key string) (uint64, bool, error) {
	var c model.ChainCursor
	err := s.db.Where("`key` = ?", key).First(&c).Error
	if err == gorm.ErrRecordNotFound {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}
	return c.LastBlock, true, nil
}

func (s *UUPSStore) SetCursor(key string, block uint64) error {
	return s.db.Save(&model.ChainCursor{Key: key, LastBlock: block}).Error
}

func (s *UUPSStore) UpsertAuctionCreated(a *model.UUPSAuction) error {
	return s.db.Save(a).Error
}

func (s *UUPSStore) UpdateAuctionBid(auctionID uint64, bidder, amount string, block uint64) error {
	return s.db.Model(&model.UUPSAuction{}).Where("auction_id = ?", auctionID).Updates(map[string]interface{}{
		"highest_bidder": bidder,
		"highest_bid":    amount,
		"updated_block":  block,
	}).Error
}

func (s *UUPSStore) MarkAuctionEnded(auctionID uint64, block uint64) error {
	return s.db.Model(&model.UUPSAuction{}).Where("auction_id = ?", auctionID).Updates(map[string]interface{}{
		"settled":       true,
		"updated_block": block,
	}).Error
}

func (s *UUPSStore) MarkAuctionCancelled(auctionID uint64, block uint64) error {
	return s.db.Model(&model.UUPSAuction{}).Where("auction_id = ?", auctionID).Updates(map[string]interface{}{
		"cancelled":     true,
		"updated_block": block,
	}).Error
}

func (s *UUPSStore) SyncAuctionFromChain(a *model.UUPSAuction) error {
	return s.db.Save(a).Error
}

func (s *UUPSStore) InsertBid(b *model.UUPSBid) (inserted bool, err error) {
	tx := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "tx_hash"}, {Name: "log_index"}},
		DoNothing: true,
	}).Create(b)
	if tx.Error != nil {
		return false, tx.Error
	}
	return tx.RowsAffected > 0, nil
}

type ListAuctionParams struct {
	Status      string // all, live, upcoming, ended, cancelled, settled
	Seller      string
	NFTContract string
	Sort        string // end_time_asc, end_time_desc, ...
	Page        int
	PageSize    int
}

func (s *UUPSStore) ListAuctions(p ListAuctionParams) ([]model.UUPSAuction, int64, error) {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 20
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	now := uint64(time.Now().Unix())

	q := s.db.Model(&model.UUPSAuction{})
	switch strings.ToLower(strings.TrimSpace(p.Status)) {
	case "", "all":
	case "live":
		q = q.Where("settled = ? AND cancelled = ? AND start_time <= ? AND end_time > ?", false, false, now, now)
	case "upcoming":
		q = q.Where("settled = ? AND cancelled = ? AND start_time > ?", false, false, now)
	case "ended":
		q = q.Where("settled = ? AND cancelled = ? AND end_time <= ?", false, false, now)
	case "cancelled":
		q = q.Where("cancelled = ?", true)
	case "settled":
		q = q.Where("settled = ?", true)
	default:
		return nil, 0, fmt.Errorf("unknown status: %s", p.Status)
	}
	if se := strings.TrimSpace(p.Seller); se != "" {
		q = q.Where("LOWER(seller) = ?", strings.ToLower(se))
	}
	if nc := strings.TrimSpace(p.NFTContract); nc != "" {
		q = q.Where("LOWER(nft_contract) = ?", strings.ToLower(nc))
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	order := "auction_id DESC"
	switch strings.ToLower(strings.TrimSpace(p.Sort)) {
	case "end_time_asc":
		order = "end_time ASC, auction_id ASC"
	case "end_time_desc":
		order = "end_time DESC, auction_id DESC"
	case "start_time_asc":
		order = "start_time ASC, auction_id ASC"
	case "start_time_desc":
		order = "start_time DESC, auction_id DESC"
	case "start_price_asc":
		order = "CAST(start_price AS DECIMAL(65,0)) ASC, auction_id ASC"
	case "start_price_desc":
		order = "CAST(start_price AS DECIMAL(65,0)) DESC, auction_id DESC"
	case "auction_id_asc":
		order = "auction_id ASC"
	case "auction_id_desc":
		order = "auction_id DESC"
	}

	offset := (p.Page - 1) * p.PageSize
	var rows []model.UUPSAuction
	err := q.Order(order).Offset(offset).Limit(p.PageSize).Find(&rows).Error
	return rows, total, err
}

func (s *UUPSStore) ListBidsByAuction(auctionID uint64, page, pageSize int) ([]model.UUPSBid, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 50
	}
	if pageSize > 200 {
		pageSize = 200
	}
	var total int64
	if err := s.db.Model(&model.UUPSBid{}).Where("auction_id = ?", auctionID).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	var rows []model.UUPSBid
	err := s.db.Where("auction_id = ?", auctionID).
		Order("block_number ASC, log_index ASC").
		Offset(offset).Limit(pageSize).Find(&rows).Error
	return rows, total, err
}

func (s *UUPSStore) Stats() (auctions, bids int64, err error) {
	err = s.db.Model(&model.UUPSAuction{}).Count(&auctions).Error
	if err != nil {
		return 0, 0, err
	}
	err = s.db.Model(&model.UUPSBid{}).Count(&bids).Error
	return auctions, bids, err
}

func (s *UUPSStore) ListStaleOpenAuctions(beforeEnd uint64, limit int) ([]model.UUPSAuction, error) {
	if limit <= 0 {
		limit = 80
	}
	var rows []model.UUPSAuction
	err := s.db.Where("settled = ? AND cancelled = ? AND end_time < ?", false, false, beforeEnd).
		Limit(limit).Find(&rows).Error
	return rows, err
}
