package model

import (
    "time"
)

type User struct {
    ID            uint      `gorm:"primaryKey"`
    WalletAddress string    `gorm:"uniqueIndex;size:42"`
    CreatedAt     time.Time
}

type Auction struct {
    ID            uint      `gorm:"primaryKey"`
    AuctionID     uint64    `gorm:"uniqueIndex;not null"`
    NFTContract   string    `gorm:"size:42;not null"`
    TokenID       uint64    `gorm:"not null"`
    StartPrice    string    `gorm:"size:78;not null"`
    HighestBid    string    `gorm:"size:78;default:0"`
    HighestBidder string    `gorm:"size:42"`
    EndTime       time.Time `gorm:"not null"`
    Ended         bool      `gorm:"default:false"`
    Creator       string    `gorm:"size:42;not null"`
    CreatedAt     time.Time
    UpdatedAt     time.Time
}

type Bid struct {
    ID          uint      `gorm:"primaryKey"`
    AuctionID   uint64    `gorm:"not null;index"`
    Bidder      string    `gorm:"size:42;not null;index"`
    Amount      string    `gorm:"size:78;not null"`
    TxHash      string    `gorm:"size:66"`
    BlockNumber uint64
    CreatedAt   time.Time
}

type SyncLog struct {
    ID          uint      `gorm:"primaryKey"`
    EventType   string    `gorm:"size:50;not null"`
    BlockNumber uint64    `gorm:"not null"`
    TxHash      string    `gorm:"size:66;not null"`
    ProcessedAt time.Time
}