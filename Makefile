.PHONY: prepare install

prepare:
	brew install pre-commit
	go install github.com/swaggo/swag/cmd/swag@latest

install:
	go mod tidy
	cd document && npm install

generate.api-spec:
	@echo "Go generate api specs..."
	go generate -run="swag" ./...

dev.document:
	cd document && npm start
