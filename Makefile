.PHONY: docker-local docker-stop \
		app-run app-build app-test \
		mockgen swag scan \
		mod-tidy mod-upgrade mod-clean-cache

# ==============================================================================
# Docker commands
docker-local:
	@echo "Starting docker compose..."
	@docker-compose -f docker-compose.local.yml up -d
	@echo "Docker compose started."

# TODO: still not working
docker-delve:
	@echo "Starting docker compose..."
	@docker-compose -f docker-compose.delve.yml up -d
	@echo "Docker compose started."

# ==============================================================================
# Application commands
app-run: swag mod-tidy
	@echo "Running application..."
	@ENV=local go run main.go
	@echo "Application running."

app-build: swag mod-tidy
	@echo "Building application..."
	@go build main.go
	@mkdir -p bin
	@mv main bin/
	@echo "Application build."

app-test:
	@echo "Testing application..."
	@go test -v ./...
	@echo "Application tested."

# ==============================================================================
# Modules support
mod-tidy:
	@echo "Tidying go modules..."
	@go mod tidy
	@echo "Go modules tidied."

mod-upgrade:
	@echo "Upgrading go modules..."
	@go get -u -t -d -v ./...
	@go mod tidy
	@echo "Go modules upgraded."

mod-clean-cache:
	@echo "Cleaning go cache..."
	@go clean -modcache
	@echo "Go cache cleaned."

# ==============================================================================
# Go migrate postgresql
migrate-up:
	@echo "Migrating database..."
	@migrate -path external/migration -database "postgresql://go-lazisnu-user:go-lazisnu-password@localhost:5432/go-lazisnu-db?sslmode=disable" --verbose up
	@echo "Database migrated."

migrate-up-1:
	@echo "Migrating database..."
	@migrate -path external/migration -database "postgresql://go-lazisnu-user:go-lazisnu-password@localhost:5432/go-lazisnu-db?sslmode=disable" --verbose up 1
	@echo "Database migrated."

migrate-down:
	@echo "Rolling back database..."
	@migrate -path external/migration -database "postgresql://go-lazisnu-user:go-lazisnu-password@localhost:5432/go-lazisnu-db?sslmode=disable" --verbose down
	@echo "Database rolled back."

migrate-down-1:
	@echo "Rolling back database..."
	@migrate -path external/migration -database "postgresql://go-lazisnu-user:go-lazisnu-password@localhost:5432/go-lazisnu-db?sslmode=disable" --verbose down 1
	@echo "Database rolled back."

# ==============================================================================
# Tools commands
# mockgen: https://github.com/uber-go/mock
# swag: https://github.com/swaggo/swag
mockgen:
	@echo "Generating mockgen..."
	@mockgen -package mockdb -destination internal/repository_generated/db/*.go github.com/rochmanramadhani/go-lazisnu/internal/repository_generated/db/mock
	@echo "mockgen generated."

swag:
	@echo "Starting swagger generating"
	@mkdir -p docs
	@swag init --propertyStrategy snakecase
	@echo "Swagger generated."

scan:
	@echo "Starting gosec scanning"
	@script/gosec.sh
	@echo "Gosec finished."
