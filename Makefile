# Makefile for 6/49 Lottery Picker (Wails + Hexagonal)

BINARY_NAME := lottery-picker
GOLANGCI_LINT_VERSION := v2.8.0 # Current as of right now

.PHONY: all build dev test lint install-tools help

all: test build

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install-fmt-tools: ## Install Go formatting tools
	@echo "Installing code formatting tools (golines, goimports, gofumpt)..."
	go install github.com/segmentio/golines@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install mvdan.cc/gofumpt@latest

fmt: ## Format Go source code
	@echo "Formatting code..."
	@golangci-lint run --fix --issues-exit-code=0

mod-upgrade: ## Upgrade Go dependencies
	go get -u ./...

dev: ## Run Wails in development mode
	wails dev

deps: ## Install Go module dependencies
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

generate: ## Generate Wails bindings
	go generate ./...

build: deps generate ## Build the production binary
	wails build

cover: ## Run tests with coverage
	go test -coverprofile=cover.out ./...
	go tool cover -func=cover.out

test: ## Run unit tests
	go test ./... -v

lint: ## Run golangci-lint
	golangci-lint run

lint-actions: ## Run actionlint to validate GitHub Actions
	actionlint

install-actionlint: ## Install actionlint
	@echo "Installing actionlint..."
	go install github.com/rhysd/actionlint/cmd/actionlint@latest

install-tools: install-fmt-tools install-actionlint ## Install development tools (golangci-lint, actionlint)
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)
	@echo "Checking for Wails..."
	@which wails > /dev/null || echo "Wails not found, installing..." || go install github.com/wailsapp/wails/v2/cmd/wails@latest

clean: ## Clean build artifacts
	rm -rf build

plantuml:
	@rm -f design/images/uml/*.png
	plantuml -DPLANTUML_LIMIT_SIZE=16384 design/uml/ -o ../images/uml/ -tpng

plantuml-svg:
	@rm -f design/images/uml/*.svg
	plantuml -DPLANTUML_LIMIT_SIZE=16384 design/uml/ -o ../images/uml/ -tsvg