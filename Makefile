.PHONY: up down
COMPOSE=docker-compose
DOCKER=docker
SQL_CONTAINER_NAME=organization-ranking-db
up:
	@${COMPOSE} up -d

down:
	@${COMPOSE} down

exec:
	@${DOCKER} exec -it ${SQL_CONTAINER_NAME} /bin/bash 

