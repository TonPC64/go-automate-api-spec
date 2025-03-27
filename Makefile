.PHONY: install test swag

install:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest

test:
	go test ./...

generate.api-spec:
	@echo "Go generate api specs..."
	go generate -run="swag" ./...

dev.document:
	cd document && npm start
