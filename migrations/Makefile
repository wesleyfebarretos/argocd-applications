MIGRATIONS_PATH := ./internal/database/migrations

create-migration:
	if [ -z "$(NAME)" ] || [ -z "$(PROJECT)" ]; then \
		echo "Usage: make create-migration NAME=<migration_name> PROJECT=<project_name>"; \
		exit 1; \
	fi

	migrate create -ext sql -dir ./$(PROJECT)/migrations/ -seq $(NAME)

.PHONY: create-migration
