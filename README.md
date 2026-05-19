# Backend Assignment

## Tech Stack

- Go
- Gin
- In-memory storage
- Mutex-based concurrency handling

---
## Architecture

The application follows a layered architecture:

- Handlers → HTTP layer
- Services → Business logic
- Storage → In-memory persistence
- Utils → Shared helpers

---

## Concurrency Design

Rate limiting uses mutex locks to prevent race conditions during concurrent requests.

---

## API Design Decisions

- Lightweight product listing endpoint
- Full product detail endpoint
- Validation at service layer
- Proper HTTP status codes

## Features

### Part 1

- Rolling 1-minute rate limiter
- Max 5 requests per user
- Concurrency safe
- Rate limit statistics

### Part 2

- Product catalog APIs
- Media URL support
- Pagination
- Optimized product listing
- SKU uniqueness validation

---

## Run Locally

```bash
go mod tidy
go run cmd/server/main.go