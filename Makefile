.PHONY: deps fmt lint test build

deps:
	go mod tidy

deps-upgrade:
	go get -u ./...
	go mod tidy

fmt:
	go fmt ./...

test:
	go test ./...

build:
	go build -o bin/mygame ./cmd/mygame

# precommit
pc: fmt test build

run:
	go run ./cmd/mygame
