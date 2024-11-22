POSTGRES_DSN=postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable
MIGRATION_DIR=db/migration

.PHONY: test
test:
	go test -v -cover -short ./...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated in coverage.html"

.PHONY: server
server:
	go run cmd/main.go

.PHONY: db-status
db-status:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(POSTGRES_DSN) goose -dir=$(MIGRATION_DIR) status

.PHONY: migrate-up
migrate-up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(POSTGRES_DSN) goose -dir=$(MIGRATION_DIR) up

.PHONY: migrate-down
migrate-down:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(POSTGRES_DSN) goose -dir=$(MIGRATION_DIR) down

.PHONY: migrate-reset
migrate-reset:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(POSTGRES_DSN) goose -dir=$(MIGRATION_DIR) reset

.PHONY: migrate-create
migrate-create:
ifndef name
	$(error name is required, use: `make migrate-create name=your_migration_name`)
endif
	GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=$(POSTGRES_DSN) goose -dir $(MIGRATION_DIR) create "$(name)" sql