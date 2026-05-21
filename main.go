package main

import (
	"backend-assignment/internal/storage"
	"backend-assignment/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	store := storage.NewMemoryStore()

	routes.RegisterRoutes(r, store)

	r.Run(":8080")
}
