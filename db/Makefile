.PHONY: help

help:
	@echo ""
	@echo "Available targets:"
	@echo "  migrate-create name=<name>   Create a new migration"
	@echo ""

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=add_new_table"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(name)

