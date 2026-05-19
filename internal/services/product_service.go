package services

import (
	"backend-assignment/internal/models"
	"backend-assignment/internal/storage"
	"backend-assignment/internal/utils"
	"errors"
	"time"

	"github.com/google/uuid"
)

func CreateProduct(store *storage.MemoryStore, req models.ProductCreateRequest) (models.Product, error) {

	if req.Name == "" || req.SKU == "" {
		return models.Product{}, errors.New("name and sku required")
	}

	if err := utils.ValidateMediaURLs(req.ImageURLs); err != nil {
		return models.Product{}, err
	}

	if err := utils.ValidateMediaURLs(req.VideoURLs); err != nil {
		return models.Product{}, err
	}

	store.ProductMutex.Lock()
	defer store.ProductMutex.Unlock()

	if _, exists := store.SKUIndex[req.SKU]; exists {
		return models.Product{}, errors.New("duplicate sku")
	}

	product := models.Product{
		ID:        uuid.New().String(),
		Name:      req.Name,
		SKU:       req.SKU,
		ImageURLs: req.ImageURLs,
		VideoURLs: req.VideoURLs,
		CreatedAt: time.Now(),
	}

	store.Products[product.ID] = product
	store.SKUIndex[product.SKU] = product.ID

	return product, nil
}
