include .env

# Взлёты

up:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

up-force:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait --force-recreate

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
	migrate -path ./xor-go/services/finances/migrations -database 'postgres://$(FINANCES_POSTGRES_USER):$(FINANCES_POSTGRES_PASSWORD)@$(FINANCES_POSTGRES_HOST_LOCAL):$(FINANCES_POSTGRES_PORT_EXTERNAL)/$(FINANCES_POSTGRES_NAME)?sslmode=disable' up

migrate-down:
	migrate -path ./xor-go/services/finances/migrations -database 'postgres://$(FINANCES_POSTGRES_USER):$(FINANCES_POSTGRES_PASSWORD)@$(FINANCES_POSTGRES_HOST_LOCAL):$(FINANCES_POSTGRES_PORT_EXTERNAL)/$(FINANCES_POSTGRES_NAME)?sslmode=disable' down 1

gen-courses:
	oapi-codegen --config xor-go/.codegen/courses-oapi-codegen.yaml oapi/courses.yaml
	oapi-codegen --config xor-go/.codegen/finances-client-product-oapi-codegen.yaml xor-go/services/finances/.codegen/product-codegen.yaml
	oapi-codegen --config xor-go/.codegen/finances-client-purchase-oapi-codegen.yaml xor-go/services/finances/.codegen/purchase-request-codegen.yaml

