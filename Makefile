include .env

create-migration:
	goose -dir ./database/postgres/migrations create example_migration sql

migrate:
	@source .env && goose -dir ./database/postgres/migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)" up

down:
	@source .env && goose -dir ./database/postgres/migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)" down

docker:
	docker-compose -f docker-compose.yaml up -d

sqlc:
	sqlc generate -f ./database/sqlc.yaml
