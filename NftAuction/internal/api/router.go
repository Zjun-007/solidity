package api

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter(h *Handler) *gin.Engine {
    r := gin.Default()
    r.GET("/health", h.HealthCheck)

    api := r.Group("/api/v1")
    {
        // 拍卖操作
        api.POST("/auction/deposit", h.DepositNFT)
        api.POST("/auction/bid", h.PlaceBid)
        api.POST("/auction/end", h.EndAuction)
        api.POST("/auction/withdraw", h.Withdraw)
        api.POST("/auction/cancel", h.CancelAuction)
        api.GET("/auction", h.GetAuctionInfo)
    }

    return r
}