ENV_DIR=env
ENV_FILE=$(ENV_DIR)/.env
DB_URL=postgres://paramet:bookmarkOverflow@localhost:5432/bookmark?sslmode=disable
MIGRATE_DIR=internal/adapter/outbound/postgresql/migrations

# docker 
dev:
	docker compose --env-file $(ENV_FILE) up -d

dev-down:
	docker compose --env-file $(ENV_FILE) down


# migrate create table
migrate-create-users:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq create_users_table

migrate-create-col:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq create_collections_table

migrate-create-book:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq create_bookmarks_table

migrate-create-tags:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq create_tags_table

migrate-create-book-tags:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq create_bookmark_tags_table

migrate-create-permissions:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq create_bookmark_permissions_table

migrate-create-all: migrate-create-users migrate-create-col migrate-create-book migrate-create-tags migrate-create-book-tags migrate-create-permissions
	@echo "All migration files created"


# connect database 
migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATE_DIR) up

migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATE_DIR) down

migrate-drop:
	migrate -database "$(DB_URL)" -path $(MIGRATE_DIR) drop -f

migrate-version:
	migrate -database "$(DB_URL)" -path $(MIGRATE_DIR) version