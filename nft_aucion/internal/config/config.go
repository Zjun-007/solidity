package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	EthRPCURL              string
	NFTAuctionProxy        string
	AuctionDeployFromBlock uint64
	IndexPollInterval      time.Duration
	IndexChunkBlocks       uint64
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	ServerPort             string
}

func Load() *Config {
	pollSec, _ := strconv.Atoi(os.Getenv("INDEX_POLL_INTERVAL_SEC"))
	if pollSec <= 0 {
		pollSec = 15
	}
	fromBlk, _ := strconv.ParseUint(strings.TrimSpace(os.Getenv("AUCTION_DEPLOY_FROM_BLOCK")), 10, 64)
	chunkBlk, _ := strconv.ParseUint(strings.TrimSpace(os.Getenv("INDEX_CHUNK_BLOCKS")), 10, 64)
	if chunkBlk <= 0 {
		chunkBlk = 2000
	}
	proxy := strings.TrimSpace(os.Getenv("NFT_AUCTION_PROXY_ADDRESS"))
	if proxy == "" {
		proxy = strings.TrimSpace(os.Getenv("NFT_AUCTION_PROXY"))
	}

	cfg := &Config{
		EthRPCURL:              strings.TrimSpace(os.Getenv("ETH_RPC_URL")),
		NFTAuctionProxy:        proxy,
		AuctionDeployFromBlock: fromBlk,
		IndexPollInterval:      time.Duration(pollSec) * time.Second,
		IndexChunkBlocks:       chunkBlk,
		DBHost:                 getenv("DB_HOST", "localhost"),
		DBPort:                 getenv("DB_PORT", "3306"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBName:                 getenv("DB_NAME", "nft_auction"),
		ServerPort:             getenv("SERVER_PORT", "8080"),
	}

	// Validate required fields
	if cfg.EthRPCURL == "" {
		log.Fatal("ETH_RPC_URL is required")
	}
	if cfg.NFTAuctionProxy == "" {
		log.Fatal("NFT_AUCTION_PROXY_ADDRESS (or NFT_AUCTION_PROXY) is required")
	}
	if cfg.DBUser == "" {
		log.Fatal("DB_USER is required")
	}
	if cfg.DBPassword == "" {
		log.Fatal("DB_PASSWORD is required")
	}

	return cfg
}

func getenv(k, def string) string {
	if v := strings.TrimSpace(os.Getenv(k)); v != "" {
		return v
	}
	return def
}
