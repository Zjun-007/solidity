package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "geth/internal/api"
    "geth/internal/blockchain"
    "geth/internal/model"
    "geth/internal/repository"
    "geth/internal/service"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system env")
    }

    // 数据库连接
    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" +
        os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
        os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // 自动迁移
    if err := db.AutoMigrate(&model.User{}, &model.Auction{}, &model.Bid{}, &model.SyncLog{}); err != nil {
        log.Fatal("AutoMigrate failed:", err)
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

    // 初始化仓储、服务
    auctionRepo := repository.NewAuctionRepository(db)
    auctionService := service.NewAuctionService(bcClient, auctionRepo)

    // 启动事件监听器（goroutine）
    eventListener, err := blockchain.NewEventListener(
        os.Getenv("ETH_WS_URL"),
        os.Getenv("CONTRACT_ADDRESS"),
        db,
    )
    if err != nil {
        log.Fatal("Failed to create event listener:", err)
    }
    go eventListener.Start()

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