# YOP Go SDK Makefile

.PHONY: help build test test-verbose test-race test-cover clean fmt lint vet deps check-deps install-tools benchmark

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
build: ## Build the project
	@echo "Building..."
	go build -v ./...

# Test targets
test: ## Run tests
	@echo "Running tests..."
	go test ./yop/...

test-verbose: ## Run tests with verbose output
	@echo "Running tests with verbose output..."
	go test -v ./yop/...

test-race: ## Run tests with race detection
	@echo "Running tests with race detection..."
	go test -race ./yop/...

test-cover: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -cover ./yop/...

test-cover-html: ## Generate HTML coverage report
	@echo "Generating HTML coverage report..."
	go test -coverprofile=coverage.out ./yop/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

benchmark: ## Run benchmarks
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./yop/...

# Code quality targets
fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./yop/...
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w ./yop/; \
	else \
		echo "goimports not found, skipping import formatting"; \
	fi

vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./yop/...

lint: install-golangci-lint ## Run golangci-lint
	@echo "Running golangci-lint..."
	golangci-lint run ./yop/...

# Dependency management
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download

tidy: ## Tidy dependencies
	@echo "Tidying dependencies..."
	go mod tidy

verify: ## Verify dependencies
	@echo "Verifying dependencies..."
	go mod verify

# Security
security: install-gosec ## Run security scan
	@echo "Running security scan..."
	gosec ./yop/...

vuln-check: install-govulncheck ## Check for vulnerabilities
	@echo "Checking for vulnerabilities..."
	govulncheck ./yop/...

# Tool installation
install-tools: install-golangci-lint install-goimports install-gosec install-govulncheck ## Install all development tools

install-golangci-lint: ## Install golangci-lint
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2)

install-goimports: ## Install goimports
	@which goimports > /dev/null || (echo "Installing goimports..." && \
		go install golang.org/x/tools/cmd/goimports@latest)

install-gosec: ## Install gosec
	@which gosec > /dev/null || (echo "Installing gosec..." && \
		go install github.com/securego/gosec/v2/cmd/gosec@latest)

install-govulncheck: ## Install govulncheck
	@which govulncheck > /dev/null || (echo "Installing govulncheck..." && \
		go install golang.org/x/vuln/cmd/govulncheck@latest)

# CI targets
ci: deps fmt vet lint test test-race security ## Run all CI checks

ci-test: ## Run tests suitable for CI
	@echo "Running CI tests..."
	go test -v -race -timeout 5m -coverprofile=coverage.out ./...

# Clean targets
clean: ## Clean build artifacts
	@echo "Cleaning..."
	go clean ./...
	rm -f coverage.out coverage.html

clean-cache: ## Clean module cache
	@echo "Cleaning module cache..."
	go clean -modcache

# Release targets
tag: ## Create a new tag (usage: make tag VERSION=v1.0.0)
	@if [ -z "$(VERSION)" ]; then echo "VERSION is required. Usage: make tag VERSION=v1.0.0"; exit 1; fi
	@echo "Creating tag $(VERSION)..."
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)

# Documentation
docs: ## Generate documentation
	@echo "Generating documentation..."
	go doc -all ./...

# Development helpers
dev-setup: install-tools deps ## Set up development environment
	@echo "Development environment setup complete!"

check: fmt vet lint test ## Run all checks

# Git hooks
install-hooks: ## Install git hooks
	@echo "Installing git hooks..."
	@mkdir -p .git/hooks
	@echo '#!/bin/sh\nmake fmt vet lint' > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "Git hooks installed!"

# Docker targets (if needed in the future)
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t yop-go-sdk .

docker-test: ## Run tests in Docker
	@echo "Running tests in Docker..."
	docker run --rm yop-go-sdk make test
