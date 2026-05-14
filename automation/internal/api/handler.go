package api

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "geth/internal/service"
)

type Handler struct {
    auctionService *service.AuctionService
}

func NewHandler(auctionService *service.AuctionService) *Handler {
    return &Handler{auctionService: auctionService}
}

// CreateAuction 创建拍卖
// POST /api/v1/auctions
func (h *Handler) CreateAuction(c *gin.Context) {
    var req service.CreateAuctionRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 从请求头获取用户钱包地址（示例：X-Wallet-Address），实际应从认证中间件获取
    creator := c.GetHeader("X-Wallet-Address")
    if creator == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "missing wallet address"})
        return
    }

    txHash, err := h.auctionService.CreateAuction(req, creator)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "auction created successfully",
        "tx_hash": txHash,
    })
}

// PlaceBid 参与竞拍
// POST /api/v1/auctions/:id/bid
func (h *Handler) PlaceBid(c *gin.Context) {
    auctionIDStr := c.Param("id")
    auctionID, err := strconv.ParseUint(auctionIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auction id"})
        return
    }

    var req service.BidRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    bidder := c.GetHeader("X-Wallet-Address")
    if bidder == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "missing wallet address"})
        return
    }

    txHash, err := h.auctionService.PlaceBid(auctionID, bidder, req.Amount)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "bid placed",
        "tx_hash": txHash,
    })
}

// EndAuction 结束拍卖
// POST /api/v1/auctions/:id/end
func (h *Handler) EndAuction(c *gin.Context) {
    auctionIDStr := c.Param("id")
    auctionID, err := strconv.ParseUint(auctionIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auction id"})
        return
    }

    txHash, err := h.auctionService.EndAuction(auctionID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "auction ended",
        "tx_hash": txHash,
    })
}

// GetAuction 获取拍卖详情
// GET /api/v1/auctions/:id
func (h *Handler) GetAuction(c *gin.Context) {
    auctionIDStr := c.Param("id")
    auctionID, err := strconv.ParseUint(auctionIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auction id"})
        return
    }

    auction, err := h.auctionService.GetAuction(auctionID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "auction not found"})
        return
    }

    c.JSON(http.StatusOK, auction)
}

// ListAuctions 获取拍卖列表
// GET /api/v1/auctions?active_only=true
func (h *Handler) ListAuctions(c *gin.Context) {
    activeOnly := c.Query("active_only") == "true"
    auctions, err := h.auctionService.ListAuctions(activeOnly)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, auctions)
}

// GetBids 获取拍卖的出价记录
// GET /api/v1/auctions/:id/bids
func (h *Handler) GetBids(c *gin.Context) {
    auctionIDStr := c.Param("id")
    auctionID, err := strconv.ParseUint(auctionIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auction id"})
        return
    }

    bids, err := h.auctionService.GetBids(auctionID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, bids)
}

// HealthCheck 健康检查
// GET /health
func (h *Handler) HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}