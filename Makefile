.PHONY: deps fmt lint test build

deps:
	go get -u ./...
	go mod tidy

fmt:
	go fmt ./...

test:
	go test -v ./...

run:
	go run ./cmd/mygame

build:
	go build -o bin/mygame ./cmd/mygame
