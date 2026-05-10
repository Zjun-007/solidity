package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"geth/internal/indexer"
	"geth/internal/service"
)

type Handler struct {
	catalog *service.Catalog
	ix      *indexer.Indexer
}

func NewHandler(catalog *service.Catalog, ix *indexer.Indexer) *Handler {
	return &Handler{catalog: catalog, ix: ix}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// GET /api/v1/indexer/status
func (h *Handler) IndexerStatus(c *gin.Context) {
	last, head, errStr := h.ix.Status()
	lag := int64(head) - int64(last)
	if lag < 0 {
		lag = 0
	}
	c.JSON(http.StatusOK, gin.H{
		"last_synced_block": last,
		"chain_head_block":  head,
		"lag":               lag,
		"last_error":        errStr,
	})
}

// GET /api/v1/auctions?status=&seller=&nft_contract=&sort=&page=&page_size=
func (h *Handler) ListAuctions(c *gin.Context) {
	var q service.AuctionListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid query parameters",
			"details": err.Error(),
		})
		return
	}
	items, total, err := h.catalog.ListAuctions(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list auctions",
			"details": err.Error(),
		})
		return
	}
	ps := q.PageSize
	if ps <= 0 {
		ps = 20
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"total": total,
		"page": q.Page,
		"page_size": ps,
	})
}

// GET /api/v1/auctions/:id/bids?page=&page_size=
func (h *Handler) ListBids(c *gin.Context) {
	idU64, err := strconv.ParseUint(strings.TrimSpace(c.Param("id")), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid auction id",
			"details": "auction id must be a valid unsigned integer",
		})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	ps, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	items, total, err := h.catalog.ListBids(idU64, page, ps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list bids",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"auction_id": idU64,
		"items": items,
		"total": total,
		"page": page,
		"page_size": ps,
	})
}

// GET /api/v1/stats
func (h *Handler) Stats(c *gin.Context) {
	a, b, err := h.catalog.Stats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get statistics",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"auction_total": a,
		"bid_total":     b,
	})
}

// GET /api/v1/wallets/:address/nfts?contracts=0x...,0x...
func (h *Handler) WalletNFTs(c *gin.Context) {
	addr := strings.TrimSpace(c.Param("address"))
	contracts := strings.TrimSpace(c.Query("contracts"))
	if contracts == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query param contracts required",
			"details": "comma-separated ERC721 addresses",
		})
		return
	}
	items, err := h.catalog.WalletNFTs(c.Request.Context(), addr, contracts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to query wallet NFTs",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"wallet": strings.ToLower(addr),
		"contracts": items,
	})
}
