package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "geth/internal/api"
    "geth/internal/blockchain"
    "geth/internal/service"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system env")
    }

    // 区块链客户端
    bcClient, err := blockchain.NewClient(
        os.Getenv("ETH_RPC_URL"),
        os.Getenv("CONTRACT_ADDRESS"),
        os.Getenv("PRIVATE_KEY"),
        11155111, // Sepolia chain ID
    )
    if err != nil {
        log.Fatal("Failed to create blockchain client:", err)
    }

    auctionService := service.NewAuctionService(bcClient)

    // 启动 HTTP 服务
    handler := api.NewHandler(auctionService)
    router := api.SetupRouter(handler)

    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server starting on :%s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}