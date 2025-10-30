# ─── Makefile ───
DOCKER_COMPOSE = docker compose
DEV_COMPOSE = docker-compose.dev.yml

# ─── Production ────────────────
up:
	$(DOCKER_COMPOSE) up --build -d

down:
	$(DOCKER_COMPOSE) down 

logs:
	$(DOCKER_COMPOSE) logs -f api

# ─── Development ───────────────
up-dev:
	$(DOCKER_COMPOSE) -f $(DEV_COMPOSE) up --build

down-dev:
	$(DOCKER_COMPOSE) -f $(DEV_COMPOSE) down 

logs-dev:
	$(DOCKER_COMPOSE) -f $(DEV_COMPOSE) logs -f api-dev

# ─── Database & Utilities ─────
db:
	docker exec -it backend_postgres psql -U postgres -d app_db

ps:
	docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

clean:
	docker system prune -af --volumes
