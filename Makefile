run:
	go run ./cmd


###### Migration Section ######
MIGRATE_CMD=go run ./cmd/migrations/migrate/migrate.go

migrate-new:
	@echo "Creating new migration:"
	@go run ./cmd/migrations/create/create.go

migrate-up:
	@echo "Migrating up..."
	@$(MIGRATE_CMD) up

migrate-down:
	@echo "Migrating down one step..."
	@$(MIGRATE_CMD) down

migrate-status:
	@echo "Migration status..."
	@$(MIGRATE_CMD) status
