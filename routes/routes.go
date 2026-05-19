package routes

import (
	"backend-assignment/internal/handlers"
	"backend-assignment/internal/storage"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, store *storage.MemoryStore) {

	requestHandler := handlers.NewRequestHandler(store)
	productHandler := handlers.NewProductHandler(store)

	r.GET("/health", handlers.HealthCheck)

	r.POST("/request", requestHandler.CreateRequest)
	r.GET("/stats", requestHandler.GetStats)

	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.ListProducts)
	r.GET("/products/:id", productHandler.GetProduct)
	r.POST("/products/:id/media", productHandler.AddMedia)
}
