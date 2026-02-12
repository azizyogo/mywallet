.PHONY: help build run dev test clean deps migrate-up migrate-down migrate-create migrate-version docker-build docker-up docker-down docker-clean docker-logs docker-restart

# Variables
BINARY_NAME=mywallet
MIGRATE_PATH=./migrations

# Load .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Default DB connection (can be overridden by .env)
DB_USER ?= mywallet_user
DB_PASS ?= mywallet_pass
DB_HOST ?= localhost
DB_PORT ?= 3306
DB_NAME ?= mywallet_db
DB_URL=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

# Default target
help:
	@echo "MyWallet - Available Commands:"
	@echo "================================"
	@echo ""
	@echo "Application:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the application"
	@echo "  make dev            - Run with auto-reload (requires air)"
	@echo "  make test           - Run tests"
	@echo "  make deps           - Download dependencies"
	@echo "  make clean          - Remove build artifacts"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build   - Build Docker image"
	@echo "  make docker-up      - Start all services with Docker Compose"
	@echo "  make docker-down    - Stop all services"
	@echo "  make docker-clean   - Stop all services and remove volumes (deletes data)"
	@echo "  make docker-logs    - View application logs"
	@echo "  make docker-restart - Restart all services"
	@echo ""
	@echo "Database Migration:"
	@echo "  make migrate-create NAME=table_name - Create new migration"
	@echo "  make migrate-up     - Apply all migrations"
	@echo "  make migrate-down   - Rollback 1 migration"
	@echo "  make migrate-version- Show current DB version"
	@echo ""

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o bin/$(BINARY_NAME) .
	@echo "✓ Build complete: bin/$(BINARY_NAME)"

# Run the application
run:
	@echo "Starting $(BINARY_NAME)..."
	@go run main.go

# Run with hot reload (requires: go install github.com/cosmtrek/air@latest)
dev:
	@echo "Starting development server with hot reload..."
	@air

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "✓ Dependencies downloaded"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean
	@echo "✓ Clean complete"

# Docker commands
docker-build:
	@echo "Building Docker image..."
	@docker-compose build
	@echo "✓ Docker image built"

docker-up:
	@echo "Starting services with Docker Compose..."
	@docker-compose up -d
	@echo "✓ Services started"
	@echo "API available at http://localhost:8080"

docker-down:
	@echo "Stopping services..."
	@docker-compose down
	@echo "✓ Services stopped"

docker-clean:
	@echo "Stopping services and removing volumes..."
	@docker-compose down -v
	@echo "⚠️  All data has been deleted"
	@echo "✓ Clean complete"

docker-logs:
	@docker-compose logs -f app

docker-restart:
	@echo "Restarting services..."
	@docker-compose restart
	@echo "✓ Services restarted"

# Database migration commands
clean:
	rm -rf bin/
	go clean
	@echo "✓ Clean complete"

# Migration: Create new migration file
migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate-create NAME=your_table_name"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(NAME)
	@echo "✓ Migration created: migrations/"

# Migration: Apply all pending migrations
migrate-up:
	migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/mywallet?charset=utf8mb4&parseTime=True&loc=Local" up
	@echo "✓ Migrations applied"

# Migration: Rollback 1 migration
migrate-down:
	migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/mywallet?charset=utf8mb4&parseTime=True&loc=Local" down 1
	@echo "✓ Migration rolled back"

# Migration: Check current version
migrate-version:
	migrate -path migrations -database "$(DB_URL)" version