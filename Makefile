PWD = ${CURDIR}
NAME = L0

# Запустить проект
.PHONY: run
run:
	go run $(PWD)/cmd/$(NAME)/

# Создать .env файл
.PHONY: local
local:
	cp .dist.env .env

# Запустить миграции
.PHONY: migrate
migrate:
	go run $(PWD)/cmd/migrations

# Запустить docker
.PHONY: docker
docker:
	docker compose up -d
