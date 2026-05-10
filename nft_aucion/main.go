package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"geth/internal/api"
	"geth/internal/config"
	"geth/internal/db"
	"geth/internal/indexer"
	"geth/internal/repository"
	"geth/internal/service"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Load optional environment variables from a .env file if present.
	// If the file is missing, proceed using the existing system environment.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	cfg := config.Load()
	if cfg.EthRPCURL == "" {
		log.Fatal("ETH_RPC_URL is required")
	}
	if cfg.NFTAuctionProxy == "" {
		log.Fatal("NFT_AUCTION_PROXY_ADDRESS (or NFT_AUCTION_PROXY) is required")
	}

	// Open the application database connection.
	gormDB, err := db.Open(cfg)
	if err != nil {
		log.Fatal("database:", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}()

	// Create a shared Ethereum client for both indexer and API queries.
	ethCli, err := ethclient.Dial(cfg.EthRPCURL)
	if err != nil {
		log.Fatal("eth client:", err)
	}
	defer ethCli.Close()

	// Initialize the blockchain event indexer with shared client.
	ix, err := indexer.New(cfg, gormDB, ethCli)
	if err != nil {
		log.Fatal("indexer:", err)
	}

	// Build the service layer and HTTP API router.
	store := repository.NewUUPSStore(gormDB)
	cat := service.NewCatalog(store, ethCli)
	handler := api.NewHandler(cat, ix)
	router := api.SetupRouter(handler)

	port := cfg.ServerPort
	log.Printf("API listening on :%s (NFT auction proxy %s)", port, cfg.NFTAuctionProxy)

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start indexer in background
	go func() {
		log.Println("Starting indexer...")
		ix.Run(ctx)
		log.Println("Indexer stopped")
	}()

	// Setup graceful shutdown on SIGINT/SIGTERM.
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		log.Println("Received shutdown signal, stopping services...")
		cancel()
	}()

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
