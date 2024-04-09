include .env

up:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

up-p:
	docker-compose up --build courses-svc finances-svc

up-r:
	docker-compose up driver-svc location-svc

up-d:
	docker-compose up --build driver-svc location-svc -d

down:
	docker-compose down driver-svc location-svc

up-obs:
	docker-compose -f ./deployments/compose.yaml up -d --build

down-obs:
	docker-compose -f ./deployments/compose.yaml down