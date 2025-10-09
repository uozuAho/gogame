.PHONY: deps fmt lint test build

deps:
	go get -u ./...
	go mod tidy

fmt:
	go fmt ./...

test:
	go test -v ./...

# precommit
pc: fmt test

build:
	go build -o bin/mygame ./cmd/mygame

run:
	go run ./cmd/mygame
