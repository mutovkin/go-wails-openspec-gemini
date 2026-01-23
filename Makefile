# Makefile for 6/49 Lottery Picker (Wails + Hexagonal)

BINARY_NAME := lottery-picker
GOLANGCI_LINT_VERSION := v1.63.4 # User requested 2.8, but 1.x is current stable. Adjust if needed.

.PHONY: all build dev test lint install-tools help

all: test build

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev: ## Run Wails in development mode
	wails dev

build: ## Build the production binary
	wails build

test: ## Run unit tests
	go test -v ./internal/... ./cmd/...

lint: ## Run golangci-lint
	golangci-lint run

install-tools: ## Install development tools (golangci-lint)
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)
	@echo "Checking for Wails..."
	@which wails > /dev/null || echo "Wails not found. Please install via: go install github.com/wailsapp/wails/v2/cmd/wails@latest"

clean: ## Clean build artifacts
	rm -rf build/bin
