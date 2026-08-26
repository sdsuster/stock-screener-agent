.PHONY: help install dev db-up db-down generate migrate test clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## Install dependencies and Encore CLI
	@echo "Installing Encore CLI..."
	npm install -g encore
	@echo "Installing Go dependencies..."
	go mod download
	@echo "Done!"

dev: ## Run services in development mode
	encore run

db-up: ## Start PostgreSQL database
	docker-compose up -d
	@echo "Waiting for database to be ready..."
	@sleep 5
	@echo "Database is ready!"

db-down: ## Stop PostgreSQL database
	docker-compose down

db-reset: ## Reset database (warning: deletes all data)
	docker-compose down -v
	docker-compose up -d

generate: ## Generate Ent code for all services
	@echo "Generating Ent code for marketdata..."
	cd marketdata && go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
	@echo "Generating Ent code for news..."
	cd news && go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
	@echo "Done!"

migrate: ## Run database migrations
	encore db migrate --schema marketdata
	encore db migrate --schema news

test: ## Run tests
	encore test ./...

test-marketdata: ## Run Market Data service tests
	encore test ./marketdata/...

test-news: ## Run News service tests
	encore test ./news/...

test-gateway: ## Run Gateway service tests
	encore test ./gateway/...

build: ## Build all services
	encore build

docker: ## Build Docker image
	encore build docker

deploy: ## Deploy to Encore Cloud
	encore deploy

clean: ## Clean generated files
	find . -name "*.gen.go" -delete
	go clean

lint: ## Run linter
	golangci-lint run

format: ## Format code
	go fmt ./...

check: format lint test ## Run format, lint, and test

.DEFAULT_GOAL := help
