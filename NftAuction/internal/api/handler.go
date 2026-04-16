package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "geth/internal/service"
)

type Handler struct {
    auctionService *service.AuctionService
}

func NewHandler(auctionService *service.AuctionService) *Handler {
    return &Handler{auctionService: auctionService}
}

// DepositNFT 卖家存入 NFT
// POST /api/v1/auction/deposit
func (h *Handler) DepositNFT(c *gin.Context) {
    txHash, err := h.auctionService.DepositNFT()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "NFT deposited", "tx_hash": txHash})
}

// PlaceBid 参与竞拍
// POST /api/v1/auction/bid
func (h *Handler) PlaceBid(c *gin.Context) {
    var req service.BidRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    txHash, err := h.auctionService.PlaceBid(req.Amount)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "bid placed", "tx_hash": txHash})
}

// EndAuction 结束拍卖
// POST /api/v1/auction/end
func (h *Handler) EndAuction(c *gin.Context) {
    txHash, err := h.auctionService.EndAuction()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "auction ended", "tx_hash": txHash})
}

// Withdraw 提取被超越的出价
// POST /api/v1/auction/withdraw
func (h *Handler) Withdraw(c *gin.Context) {
    txHash, err := h.auctionService.Withdraw()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "withdrawn", "tx_hash": txHash})
}

// CancelAuction 取消拍卖
// POST /api/v1/auction/cancel
func (h *Handler) CancelAuction(c *gin.Context) {
    txHash, err := h.auctionService.CancelAuction()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "auction canceled", "tx_hash": txHash})
}

// GetAuctionInfo 获取当前拍卖信息
// GET /api/v1/auction
func (h *Handler) GetAuctionInfo(c *gin.Context) {
    info, err := h.auctionService.GetAuctionInfo(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, info)
}

// HealthCheck 健康检查
func (h *Handler) HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}