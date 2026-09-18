include .env
export

export PROJECT_ROOT=$(shell pwd)
env-up:
	 	@docker compose up  -d todoapp-postgres 

env-down:
		@docker compose down todoapp-postgres 
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq \
		"$(seq)"
migrate-up:
	 docker compose run --rm todoapp-postgres-migrate \
	 	-path /migrations \
	 	-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
	 	up
migrate-down:
	docker compose run --rm todoapp-postgres-migrate \
	 	-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		down
migrate-force:
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		force 1		
env-port-forward:
	@docker compose up -d port-forwarder
	
env-port-close:
	@docker compose down port-forwarder 
