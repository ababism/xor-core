include .env

up:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

up-all:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait
	docker-compose -f ./deployments/compose.yaml up -d --build

up-p:
	docker-compose up --build courses-svc finances-svc

up-c:
	docker-compose up --build courses-svc

up-d:
	docker-compose up --build courses-svc finances-svc -d

down:
	docker-compose down courses-svc finances-svc

down-c:
	docker-compose down courses-svc

up-obs:
	docker-compose -f ./deployments/compose.yaml up -d --build

down-obs:
	docker-compose -f ./deployments/compose.yaml down