package handlers

import (
	"backend-assignment/internal/models"
	"backend-assignment/internal/services"
	"backend-assignment/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Store *storage.MemoryStore
}

func NewRequestHandler(store *storage.MemoryStore) *RequestHandler {
	return &RequestHandler{
		Store: store,
	}
}

func (h *RequestHandler) CreateRequest(c *gin.Context) {

	var req models.RequestPayload

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}

	if req.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id required",
		})
		return
	}

	allowed, accepted, rejected := services.AllowRequest(h.Store, req.UserID)

	if !allowed {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error":             "rate limit exceeded",
			"accepted_requests": accepted,
			"rejected_requests": rejected,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":           "request accepted",
		"accepted_requests": accepted,
	})
}

func (h *RequestHandler) GetStats(c *gin.Context) {

	h.Store.RateMutex.Lock()
	defer h.Store.RateMutex.Unlock()

	response := map[string]interface{}{}

	for userID, data := range h.Store.RateData {

		response[userID] = models.UserStats{
			Accepted: len(data.Timestamps),
			Rejected: data.Rejected,
		}
	}

	c.JSON(http.StatusOK, response)
}
