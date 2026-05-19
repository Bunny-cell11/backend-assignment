````md
# Backend Assignment

## Overview

This project is a backend system built in Go that implements:

1. A concurrency-safe rolling window rate limiter
2. Product catalog management APIs

The implementation focuses on:
- clean architecture
- concurrency safety
- validation
- pagination
- optimized API responses
- Docker support
- testing
- maintainability

---

# Tech Stack

| Technology | Purpose |
|---|---|
| Go | Core backend language |
| Gin | HTTP framework |
| Mutexes | Concurrency safety |
| Docker | Containerization |
| In-Memory Storage | Lightweight persistence |

---

# Features

## Part 1 — Rolling Window Rate Limiter

### Functionality

- Rolling 1-minute window
- Maximum 5 requests per user
- Concurrency-safe implementation
- Per-user statistics tracking
- Proper HTTP status handling

### Endpoints

#### Create Request

```http
POST /request
````

#### Get Statistics

```http
GET /stats
```

---

## Part 2 — Product Catalog APIs

### Functionality

* Create product
* List products
* Fetch product details
* Add media to products
* SKU uniqueness validation
* Pagination support
* Optimized lightweight product summaries

### Endpoints

#### Create Product

```http
POST /products
```

#### Get Products

```http
GET /products
```

#### Get Product Details

```http
GET /products/:id
```

#### Add Product Media

```http
POST /products/:id/media
```

---

# Project Structure

```text
backend-assignment/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   │   ├── product_handler.go
│   │   ├── request_handler.go
│   │   └── health_handler.go
│   │
│   ├── middleware/
│   │   └── logger.go
│   │
│   ├── models/
│   │   ├── product.go
│   │   └── request.go
│   │
│   ├── services/
│   │   ├── product_service.go
│   │   └── rate_limiter_service.go
│   │
│   ├── storage/
│   │   └── memory.go
│   │
│   └── utils/
│       ├── pagination.go
│       └── validation.go
│
├── routes/
│   └── routes.go
│
├── tests/
│   ├── product_test.go
│   └── request_test.go
│
├── Dockerfile
├── Makefile
├── README.md
├── go.mod
└── postman_collection.json
```

---

# Architecture

The application follows a layered architecture:

## Handlers Layer

Responsible for:

* request parsing
* response formatting
* HTTP status handling

## Services Layer

Responsible for:

* business logic
* validation rules
* rate limiting logic
* product creation logic

## Storage Layer

Responsible for:

* shared in-memory data structures
* concurrency-safe access

## Utility Layer

Responsible for:

* pagination helpers
* validation helpers

---

# Concurrency Design

The rate limiter uses mutex-based synchronization to ensure thread-safe access to shared request counters.

This prevents:

* race conditions
* inconsistent request tracking
* concurrent mutation issues

Mutexes are also used during:

* product creation
* SKU uniqueness checks
* statistics updates

---

# API Design Decisions

* Lightweight product listing endpoint
* Full product detail endpoint
* Validation at service layer
* Proper HTTP status codes
* Optimized response payloads
* Clean separation of concerns

---

# Validation Rules

## Product Validation

| Validation     | Rule     |
| -------------- | -------- |
| Product Name   | Required |
| SKU            | Required |
| Duplicate SKU  | Rejected |
| URL Validation | Required |
| Max URLs       | 20       |
| Max URL Length | 2048     |

---

## Request Validation

| Validation   | Rule     |
| ------------ | -------- |
| user_id      | Required |
| Invalid JSON | Rejected |

---

# Pagination

The products listing endpoint supports:

* limit
* offset

Example:

```http
GET /products?limit=20&offset=0
```

## Pagination Constraints

| Parameter       | Rule           |
| --------------- | -------------- |
| Default Limit   | 20             |
| Maximum Limit   | 100            |
| Negative Offset | Converted to 0 |

---

# Performance Optimization

The `GET /products` endpoint intentionally returns lightweight product summaries instead of full media arrays.

This improves:

* response time
* serialization performance
* memory efficiency
* bandwidth usage

Detailed media data is returned only in:

```http
GET /products/:id
```

---

# HTTP Status Codes

| Status Code | Meaning             |
| ----------- | ------------------- |
| 200         | Success             |
| 201         | Resource Created    |
| 400         | Invalid Request     |
| 404         | Resource Not Found  |
| 409         | Duplicate SKU       |
| 429         | Rate Limit Exceeded |

---

# Running Locally

## Prerequisites

* Go 1.24+
* Git
* Docker (optional)

---

## Install Dependencies

```bash
go mod tidy
```

---

## Run Application

```bash
go run cmd/server/main.go
```

Server runs on:

```text
http://localhost:8080
```

---

# Example API Requests

## Health Check

```bash
curl http://localhost:8080/health
```

---

## Create Request

```bash
curl -X POST http://localhost:8080/request \
-H "Content-Type: application/json" \
-d '{
  "user_id":"user1",
  "payload":{"message":"hello"}
}'
```

---

## Create Product

```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{
  "name":"Widget A",
  "sku":"SKU-001",
  "image_urls":[
    "https://cdn.example.com/image1.jpg"
  ],
  "video_urls":[
    "https://cdn.example.com/video1.mp4"
  ]
}'
```

---

## Get Products

```bash
curl http://localhost:8080/products
```

---

# Running Tests

```bash
go test ./...
```

---

# Format Code

```bash
go fmt ./...
```

---

# Verify Build

```bash
go build ./...
```

---

# Docker Support

## Build Docker Image

```bash
docker build -t backend-assignment .
```

---

## Run Docker Container

```bash
docker run -p 8080:8080 backend-assignment
```

---

# Design Decisions

## Why In-Memory Storage?

The assignment requirements can be fully satisfied without introducing database complexity.

This approach:

* simplifies setup
* improves development speed
* keeps focus on backend fundamentals

---

## Why Mutex Locking?

Mutexes provide safe concurrent access to shared memory structures during:

* rate limiting
* product creation
* statistics tracking

---

## Why Lightweight Product Listing?

Returning summarized product data improves:

* scalability
* response size
* serialization efficiency

---

# Assumptions

* Media files are externally hosted
* Authentication is not required
* Single-instance deployment is acceptable
* In-memory persistence is sufficient for assignment scope

---

# Tradeoffs

* In-memory storage used instead of persistent database
* Mutex-based synchronization used instead of distributed locking
* Simplicity prioritized over production-scale infrastructure

---

# Known Limitations

* Data resets on server restart
* No persistent database
* No authentication/authorization
* No distributed rate limiting
* Single-instance only

---

# Future Improvements

Potential production-level improvements:

* PostgreSQL integration
* Redis-backed distributed rate limiter
* JWT authentication
* Structured logging
* OpenTelemetry tracing
* Swagger/OpenAPI documentation
* Kubernetes deployment
* CI/CD pipelines
* Metrics and monitoring

---

# AI Assistance Disclosure

AI tools were used during development for:

* brainstorming implementation approaches
* reviewing API design ideas
* improving documentation quality
* discussing validation and concurrency strategies

All code was manually integrated, tested, debugged, and verified before submission.

---

# Assignment Completion Status

Successfully implemented:

* concurrency-safe rate limiter
* product APIs
* pagination
* validations
* unit tests
* Docker support
* README documentation
* optimized API responses
* proper HTTP status handling

```
```
