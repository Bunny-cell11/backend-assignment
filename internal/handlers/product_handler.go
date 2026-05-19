package handlers

import (
	"backend-assignment/internal/models"
	"backend-assignment/internal/services"
	"backend-assignment/internal/storage"
	"backend-assignment/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Store *storage.MemoryStore
}

func NewProductHandler(store *storage.MemoryStore) *ProductHandler {
	return &ProductHandler{
		Store: store,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {

	var req models.ProductCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}

	product, err := services.CreateProduct(h.Store, req)

	if err != nil {

		if err.Error() == "duplicate sku" {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	limit, offset = utils.GetPagination(limit, offset)

	h.Store.ProductMutex.RLock()
	defer h.Store.ProductMutex.RUnlock()

	summaries := []models.ProductSummary{}

	count := 0

	for _, product := range h.Store.Products {

		if count < offset {
			count++
			continue
		}

		if len(summaries) >= limit {
			break
		}

		summary := models.ProductSummary{
			ID:         product.ID,
			Name:       product.Name,
			SKU:        product.SKU,
			ImageCount: len(product.ImageURLs),
			VideoCount: len(product.VideoURLs),
			CreatedAt:  product.CreatedAt,
		}

		if len(product.ImageURLs) > 0 {
			summary.ThumbnailURL = product.ImageURLs[0]
		}

		summaries = append(summaries, summary)
	}

	c.JSON(http.StatusOK, gin.H{
		"limit":    limit,
		"offset":   offset,
		"products": summaries,
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {

	id := c.Param("id")

	h.Store.ProductMutex.RLock()
	defer h.Store.ProductMutex.RUnlock()

	product, exists := h.Store.Products[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) AddMedia(c *gin.Context) {

	id := c.Param("id")

	var req models.MediaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}

	if len(req.ImageURLs) == 0 && len(req.VideoURLs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "at least one media array required",
		})
		return
	}

	if err := utils.ValidateMediaURLs(req.ImageURLs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.ValidateMediaURLs(req.VideoURLs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Store.ProductMutex.Lock()
	defer h.Store.ProductMutex.Unlock()

	product, exists := h.Store.Products[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	product.ImageURLs = append(product.ImageURLs, req.ImageURLs...)
	product.VideoURLs = append(product.VideoURLs, req.VideoURLs...)

	h.Store.Products[id] = product

	c.JSON(http.StatusOK, product)
}
