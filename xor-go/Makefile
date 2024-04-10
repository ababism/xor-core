include .env

# Взлёты

up:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

up-all:
	docker-compose -f ./deployments/compose.yaml up -d --build
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

up-courses:
	docker-compose up --build courses-svc

up-finances:
	docker-compose up --build finances-svc

up-cf:
	docker-compose up --build courses-svc

up-d:
	docker-compose up --build courses-svc finances-svc -d

# Падения

down:
	docker-compose down courses-svc finances-svc

down-c:
	docker-compose down courses-svc

# Observability

up-obs:
	docker-compose -f ./deployments/compose.yaml up -d --build

down-obs:
	docker-compose -f ./deployments/compose.yaml down

# Миграции

migrate-up:
	migrate -path ./xor-go/services/finances/migrations -database 'postgres://$(FINANCES_POSTGRES_USER):$(FINANCES_POSTGRES_PASSWORD)@$(FINANCES_POSTGRES_HOST):$(FINANCES_POSTGRES_PORT)/$(FINANCES_POSTGRES_NAME)?sslmode=disable' up

migrate-down:
	migrate -path ./xor-go/services/finances/migrations -database 'postgres://$(FINANCES_POSTGRES_USER):$(FINANCES_POSTGRES_PASSWORD)@$(FINANCES_POSTGRES_HOST):$(FINANCES_POSTGRES_PORT)/$(FINANCES_POSTGRES_NAME)?sslmode=disable' down 1
