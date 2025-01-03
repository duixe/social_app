include .env
MIGRATIONS_PATH= ./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) up

.PHONY: migrate-up
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/migrate/seeder/main.go