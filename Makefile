run:
	go run cmd/server/main.go

build:
	go build -o server ./cmd/server

test:
	go test ./...

docker-build:
	docker build -t backend-assignment .

docker-run:
	docker run -p 8080:8080 backend-assignment