.PHONY: help
help:
	@echo ""
	@echo "Available targets:"
	@echo "  migrate-create name=<name>   Create a new migration"
	@echo "  migrate-up                  Run all up migrations (local DB)"
	@echo "  migrate-down                Roll back one migration"
	@echo ""

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=add_new_table"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migrate-up
migrate-up:
	migrate \
		-path migrations \
		-database "postgres://gainline:gainline@localhost:5432/gainline?sslmode=disable" \
		up

.PHONY: migrate-down
migrate-down:
	migrate \
		-path migrations \
		-database "postgres://gainline:gainline@localhost:5432/gainline?sslmode=disable" \
		down 1
