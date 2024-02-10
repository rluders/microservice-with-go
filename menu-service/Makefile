DATABASE_USERNAME ?= postgres
DATABASE_PASSWORD ?= password
DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_NAME ?= microservice_menu
DATABASE_SSL ?= disable

lint:
	golangci-lint run

migrate-up:
	migrate \
	-path ./database/migrations \
	-database "postgres://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=$(DATABASE_SSL)" \
	up

migrate-down:
	migrate \
	-path ./database/migrations \
	-database "postgres://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=$(DATABASE_SSL)" \
	down

test:
	go test ./...

run:
	go run cmd/service/main.go --config ./config/config.yaml