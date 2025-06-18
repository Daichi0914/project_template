# ビルドしてupする
build:
	docker compose build

# 再ビルドしてupする
rebuild:
	docker compose down -v
	docker compose build --no-cache

# 全コンテナを起動する
up:
	docker compose up -d

# devコンテナを起動する
up-dev:
	docker compose up -d $(shell docker compose config --services | grep dev);

# testコンテナを起動する
up-test:
	docker compose up -d $(shell docker compose config --services | grep test);

# コンテナを再起動する
reup:
	docker compose down -v
	docker compose up -d

# devコンテナを再起動する
reup-dev:
	docker compose down -v
	docker compose up -d $(shell docker compose config --services | grep dev);

# testコンテナを再起動する
reup-test:
	docker compose down -v
	docker compose up -d $(shell docker compose config --services | grep test);

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
	set -a && source .env.dev && set +a && docker compose exec $$MYSQL_HOST mysql -u$$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE

# テストデータベースに接続する
exec_test_db:
	set -a && source .env.test && set +a && docker compose exec $$MYSQL_HOST mysql -u$$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE

# サンプルデータを作成する場合は`sample.sqlに必要なデータを記述した上で`make seed`を実行する
seed_dev:
	set -a && source .env.dev && set +a && docker compose exec -T $$MYSQL_HOST mysql -u$$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE < ./backend/infrastructure/db/seed/sample.sql

seed_test:
	set -a && source .env.test && set +a && docker compose exec -T $$MYSQL_HOST mysql -u$$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE < ./backend/infrastructure/db/seed/sample.sql

# backendのテストを実行する
test_be:
	docker-compose exec backend_test go test -v ./...

# frontendのテストを実行する
test_fe:
	docker-compose exec frontend_test npm test
