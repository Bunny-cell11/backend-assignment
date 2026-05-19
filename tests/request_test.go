package tests

import (
	"backend-assignment/internal/services"
	"backend-assignment/internal/storage"
	"testing"
)

func TestRateLimiter(t *testing.T) {

	store := storage.NewMemoryStore()

	for i := 0; i < 5; i++ {

		allowed, _, _ := services.AllowRequest(store, "user1")

		if !allowed {
			t.Fail()
		}
	}

	allowed, _, _ := services.AllowRequest(store, "user1")

	if allowed {
		t.Fail()
	}
}
