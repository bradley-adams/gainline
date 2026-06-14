.PHONY: help build up rebuild down clean restart logs \
        db-up db-stop migrate db-reset \
        redis-up redis-stop redis-cli

# Default target
help:
	@echo ""
	@echo "Gainline Commands"
	@echo "------------------"
	@echo "make build        Build containers"
	@echo "make up           Start all services"
	@echo "make rebuild      Rebuild and start services"
	@echo "make down         Stop all services"
	@echo "make clean        Remove all containers + volumes"
	@echo "make restart      Restart everything"
	@echo "make logs         Tail logs"
	@echo ""
	@echo "Database"
	@echo "make db-up        Start DB only"
	@echo "make db-stop      Stop DB only"
	@echo "make migrate      Run migrations"
	@echo "make db-reset     Remove DB volume"
	@echo ""
	@echo "Redis"
	@echo "make redis-up     Start Redis only"
	@echo "make redis-stop   Stop Redis only"
	@echo "make redis-cli    Open a redis-cli session"
	@echo ""

# Build containers
build:
	docker compose build

# Build + start all services
up:
	docker compose up -d

# Force rebuild + start all services
rebuild:
	docker compose up -d --build

# Stop all services
down:
	docker compose down

# Clean everything
clean:
	docker compose down -v --remove-orphans

# Restart everything
restart: down up

# Tail logs
logs:
	docker compose logs -f

# Start DB only
db-up:
	docker compose up -d gainline-db

# Run migrations
migrate:
	docker compose run --rm gainline-migrate

# Stop DB container
db-stop:
	docker compose stop gainline-db

# Remove DB container + volume, then recreate + migrate
db-reset:
	docker compose down
	docker volume rm gainline-data || true
	docker compose up -d gainline-db
	docker compose run --rm gainline-migrate

# Start Redis only
redis-up:
	docker compose up -d gainline-redis

# Stop Redis container
redis-stop:
	docker compose stop gainline-redis

# Open a redis-cli session against the running container
redis-cli:
	docker compose exec gainline-redis redis-cli