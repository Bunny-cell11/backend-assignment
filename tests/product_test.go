package tests

import (
	"backend-assignment/internal/models"
	"backend-assignment/internal/services"
	"backend-assignment/internal/storage"
	"testing"
)

func TestCreateProduct(t *testing.T) {

	store := storage.NewMemoryStore()

	req := models.ProductCreateRequest{
		Name: "Widget A",
		SKU:  "SKU-001",
	}

	product, err := services.CreateProduct(store, req)

	if err != nil {
		t.Fail()
	}

	if product.SKU != "SKU-001" {
		t.Fail()
	}
}
