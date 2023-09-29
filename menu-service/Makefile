lint:
	golangci-lint run

migrate:
	migrate -path ./database/migrations -database "postgres://postgres:bob@localhost/microservice_menu?sslmode=disable" up