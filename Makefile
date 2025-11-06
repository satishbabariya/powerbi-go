.PHONY: help build test lint fmt clean install examples

# Default target
.DEFAULT_GOAL := help

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the SDK
	@echo "Building..."
	@go build -v ./...

test: ## Run tests
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

test-short: ## Run short tests
	@echo "Running short tests..."
	@go test -short -v ./...

lint: ## Run linter
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@gofmt -s -w .

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -f coverage.out coverage.html
	@go clean -cache -testcache -modcache

install: ## Install dependencies
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

examples: ## Run all examples (requires environment variables)
	@echo "Running examples..."
	@cd examples/basic && go run main.go
	@echo ""
	@echo "To run other examples, set the required environment variables and run:"
	@echo "  cd examples/embed-token && go run main.go"
	@echo "  cd examples/admin-operations && go run main.go"
	@echo "  cd examples/dataset-refresh && go run main.go"
	@echo "  cd examples/workspace-management && go run main.go"

check: fmt lint test ## Run all checks (format, lint, test)

coverage: test ## Generate coverage report
	@go tool cover -func=coverage.out

deps: ## Update dependencies
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy

verify: ## Verify dependencies
	@echo "Verifying dependencies..."
	@go mod verify

doc: ## Generate and open documentation
	@echo "Generating documentation..."
	@godoc -http=:6060 &
	@sleep 2
	@open http://localhost:6060/pkg/github.com/satishbabariya/powerbi-go/

