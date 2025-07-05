# Makefile for multi-user bookmark API project
DC = docker compose


# Dir names
DIR_SRC = env
ENV_DEV = --env-file .env.dev

dev:
	$(DC) $(DIR_SRC)/${ENV_DEV} up --build

dev-down:
	$(DC) $(DIR_SRC)/${ENV_DEV} down

