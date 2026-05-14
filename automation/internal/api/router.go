package api

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter(h *Handler) *gin.Engine {
    r := gin.Default()

    // 健康检查
    r.GET("/health", h.HealthCheck)

    api := r.Group("/api/v1")
    {
        // 拍卖相关
        api.POST("/auctions", h.CreateAuction)
        api.GET("/auctions", h.ListAuctions)
        api.GET("/auctions/:id", h.GetAuction)
        api.POST("/auctions/:id/bid", h.PlaceBid)
        api.POST("/auctions/:id/end", h.EndAuction)
        api.GET("/auctions/:id/bids", h.GetBids)
    }

    return r
}