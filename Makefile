# ===== Variables =====
APP_NAME      := imperio
CONFIG_FILE   := config.yaml
SCHEMA_FILE   := schema.yaml
DOCKER_IMAGE  := imperio-app
DOCKER_TAG    := latest
DOCKER_COMPOSE := docker-compose
GO_FILES      := $(shell find . -type f -name '*.go')
BIN_PATH      := ./bin/$(APP_NAME)

# ===== Targets =====

.PHONY: all build run test lint fmt docker docker-run clean help

## Default target
all: build

## Build the Go binary
build:
	@echo "🔨 Building $(APP_NAME)..."
	go build -o $(BIN_PATH) ./cmd

## Run the binary locally
run: build
	@echo "🚀 Running $(APP_NAME)..."
	$(BIN_PATH) --config=$(CONFIG_FILE)

## Run tests
test:
	@echo "🧪 Running tests..."
	go test ./... -v

## Lint the code
lint:
	@echo "🔍 Linting..."
	golangci-lint run

## Format the code
fmt:
	@echo "🎨 Formatting code..."
	go fmt ./...

## Build Docker image
docker:
	@echo "🐳 Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

## Run using Docker Compose (default = postgres)
docker-run:
	@echo "🚀 Starting using docker-compose..."
	$(DOCKER_COMPOSE) up --build

## Run using PostgreSQL (override config before use)
docker-run-postgres:
	@echo "🚀 Starting with PostgreSQL backend..."
	$(DOCKER_COMPOSE) -f docker-compose.yml -f docker-compose.postgres.yml up --build

## Stop containers
docker-stop:
	@echo "🛑 Stopping docker-compose..."
	$(DOCKER_COMPOSE) down

## Clean build artifacts
clean:
	@echo "🧹 Cleaning up..."
	rm -f $(BIN_PATH)

## Show all available commands
help:
	@echo "💡 Usage:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
