package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/health", h.Health)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/indexer/status", h.IndexerStatus)
		v1.GET("/auctions", h.ListAuctions)
		v1.GET("/auctions/:id/bids", h.ListBids)
		v1.GET("/stats", h.Stats)
		v1.GET("/wallets/:address/nfts", h.WalletNFTs)
	}
	return r
}
