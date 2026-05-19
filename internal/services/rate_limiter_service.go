package services

import (
	"backend-assignment/internal/storage"
	"time"
)

const Limit = 5
const Window = time.Minute

func AllowRequest(store *storage.MemoryStore, userID string) (bool, int, int) {

	store.RateMutex.Lock()
	defer store.RateMutex.Unlock()

	now := time.Now()

	data, exists := store.RateData[userID]

	if !exists {
		data = &storage.UserRateData{}
		store.RateData[userID] = data
	}

	valid := []time.Time{}

	for _, t := range data.Timestamps {

		if now.Sub(t) < Window {
			valid = append(valid, t)
		}
	}

	data.Timestamps = valid

	if len(valid) >= Limit {
		data.Rejected++
		return false, len(valid), data.Rejected
	}

	data.Timestamps = append(data.Timestamps, now)

	return true, len(data.Timestamps), data.Rejected
}
