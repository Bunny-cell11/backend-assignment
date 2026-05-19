package storage

import (
	"backend-assignment/internal/models"
	"sync"
	"time"
)

type UserRateData struct {
	Timestamps []time.Time
	Rejected   int
}

type MemoryStore struct {
	ProductMutex sync.RWMutex
	RateMutex    sync.Mutex

	Products map[string]models.Product
	SKUIndex map[string]string
	RateData map[string]*UserRateData
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Products: make(map[string]models.Product),
		SKUIndex: make(map[string]string),
		RateData: make(map[string]*UserRateData),
	}
}
