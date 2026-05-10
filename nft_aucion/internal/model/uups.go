package model

import "time"

const CursorKeyNFTAuction = "nft_auction_uups_logs"

type ChainCursor struct {
	Key       string `gorm:"primaryKey;size:64"`
	LastBlock uint64 `gorm:"not null"`
}

func (ChainCursor) TableName() string { return "chain_cursors" }

// UUPSAuction 镜像链上 NFTAuctionUUPS 拍卖状态（由索引器维护）
type UUPSAuction struct {
	AuctionID     uint64 `gorm:"primaryKey;column:auction_id"`
	Seller        string `gorm:"size:42;not null;index"`
	NFTContract   string `gorm:"size:42;not null;index:idx_nft"`
	TokenID       string `gorm:"size:78;not null"`
	StartPrice    string `gorm:"size:78;not null"`
	StartTime     uint64 `gorm:"not null"`
	EndTime       uint64 `gorm:"not null;index:idx_end"`
	HighestBidder string `gorm:"size:42"`
	HighestBid    string `gorm:"size:78"`
	Settled       bool   `gorm:"not null;default:false;index"`
	Cancelled     bool   `gorm:"not null;default:false;index"`
	CreatedBlock  uint64
	UpdatedBlock  uint64
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (UUPSAuction) TableName() string { return "uups_auctions" }

type UUPSBid struct {
	ID          uint   `gorm:"primaryKey"`
	AuctionID   uint64 `gorm:"not null;index:idx_auction_time,priority:1"`
	Bidder      string `gorm:"size:42;not null;index"`
	Amount      string `gorm:"size:78;not null"`
	BlockNumber uint64 `gorm:"not null"`
	TxHash      string `gorm:"size:66;not null;uniqueIndex:uk_txlog,priority:1"`
	LogIndex    uint   `gorm:"not null;uniqueIndex:uk_txlog,priority:2"`
	CreatedAt   time.Time
}

func (UUPSBid) TableName() string { return "uups_bids" }
