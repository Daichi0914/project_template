# ビルドしてupする
build:
	docker compose build
	docker compose up -d

# 再ビルドしてupする
rebuild:
	docker compose down -v
	docker compose build --no-cache
	docker compose up -d

# 全コンテナを起動する
up:
	docker compose up -d

# コンテナを再起動する
reup:
	docker compose down -v
	docker compose up -d

# コンテナのログを表示する
logs:
	docker compose logs -f

# コンテナ、ボリューム、ネットワーク、イメージを削除する
down:
	docker compose down -v

# フロントエンドコンテナを起動する
frontend:
	docker compose up -d frontend

# バックエンドコンテナを起動する
backend:
	docker compose up -d backend

# データベースコンテナを起動する
db:
	docker compose up -d db

# コンテナ、ボリューム、ネットワーク、イメージを完全に削除する
prune:
	docker system prune -f
	docker volume prune -f
	docker network prune -f
	docker image prune -f
	docker container prune -f

# データベースに接続する
exec_db:
	set -a && source .env && set +a && docker compose exec db mysql -u$$DB_USER -p$$DB_PASSWORD $$DB_NAME

# サンプルデータを作成する場合は`sample.sqlに必要なデータを記述した上で`make seed`を実行する
seed:
	set -a && source .env && set +a && docker compose exec -T db mysql -u$$DB_USER -p$$DB_PASSWORD $$DB_NAME < ./backend/infrastructure/db/seed/sample.sql

# backendのテストを実行する
test_be:
	docker-compose exec backend go test -v ./...

# frontendのテストを実行する
test_fe:
	docker-compose exec frontend npm test
