# ─── Makefile ───
COMPOSE = docker compose

# ─── Run Modes ───
dev:
	$(COMPOSE) --profile dev up --build

prod:
	$(COMPOSE) --profile prod up --build -d

down:
	$(COMPOSE) down -v

logs:
	$(COMPOSE) logs -f api-dev

# ─── DB Utilities ───
db:
	docker exec -it backend_postgres psql -U postgres -d app_db

clean:
	docker system prune -af --volumes
