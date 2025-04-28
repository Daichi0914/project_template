build:
	docker compose build

up:
	docker compose up -d

reup:
	docker compose down -v
	docker compose up -d

logs:
	docker compose logs -f

down:
	docker compose down -v

rebuild:
	docker compose down -v
	docker compose build --no-cache
	docker compose up -d

frontend:
	docker compose up -d frontend

backend:
	docker compose up -d backend

db:
	docker compose up -d db

prune:
	docker system prune -f
	docker volume prune -f
	docker network prune -f
	docker image prune -f
	docker container prune -f

exec_db:
	set -a && source .env && set +a && docker compose exec db mysql -u$$DB_USER -p$$DB_PASSWORD $$DB_NAME

seed:
	set -a && source .env && set +a && docker compose exec -T db mysql -u$$DB_USER -p$$DB_PASSWORD $$DB_NAME < ./backend/infrastructure/db/seed/sample.sql
