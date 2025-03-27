.PHONY: prepare install

prepare:
	@echo "Install pre-commit and swag..."
	brew install pre-commit
	go install github.com/swaggo/swag/cmd/swag@latest

install:
	@echo "Install go and node dependencies..."
	go mod download
	cd document && npm install

re-hooks:
	@echo "Re-install pre-commit hooks..."
	pre-commit install -f --install-hooks

generate.api-spec:
	@echo "Go generate api specs..."
	swag fmt
	go generate -run="swag" ./...

dev.document:
	@make generate.api-spec
	@echo "Start document server..."
	cd document && npm start
