.PHONY: install test build swag

install:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest

test:
	go test ./...

build:
	go build -o bin/go-automate-api-spec

swag:
	$(MAKE) -C cmd/echo swag
	$(MAKE) -C cmd/gin swag