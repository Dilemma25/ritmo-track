#################################################
# Управление контейнерами для разработки локально

up-dev:
	@docker compose -f docker-compose-dev.yml up --build

down-dev:
	@docker compose -f docker-compose-dev.yml down

stop-dev:
	@docker compose -f docker-compose-dev.yml stop

start-dev:
	@docekr compose -f docker-compose-dev.yml start

restart-dev:
	@docker compose -f docker-compose-dev.yml restart

#################################################
# Управление миграциями для разработки

upgrade-migration:
	@docker exec ritmotrack-backend go run /app/main.go migration up

down-migration:
	@docker exec ritmotrack-backend go run /app/main.go migration down