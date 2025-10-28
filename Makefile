# ────────────── Backend-Base Makefile ──────────────

APP_NAME = go_backend
DOCKER_COMPOSE = docker compose
GO = go

.PHONY: up down build logs api db test migrate fmt tidy

# Start all containers
up:
	$(DOCKER_COMPOSE) up --build

# Stop and remove containers + volumes
down:
	$(DOCKER_COMPOSE) down -v

# Build Go binary locally (optional)
build:
	$(GO) build -o bin/server ./cmd/app

# Follow API logs
logs:
	$(DOCKER_COMPOSE) logs -f api

# Run API locally (without Docker)
api:
	$(GO) run ./cmd/app/main.go

# Open PostgreSQL shell
db:
	docker exec -it backend_postgres psql -U postgres -d app_db

# Run tests
test:
	$(GO) test ./... -v

# Format + tidy
fmt:
	$(GO) fmt ./...
tidy:
	$(GO) mod tidy
