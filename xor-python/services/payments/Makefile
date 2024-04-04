
up:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait


up-d:
	docker compose --file ./deploy/docker-compose.yml --env-file ./.env up -d --build --wait

up-df:
	docker build --file ./deploy/Dockerfile .


up-p:
	poetry run python -m payments

