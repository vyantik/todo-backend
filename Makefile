include .env
export

build:
	docker-compose build todo-app

run:
	docker-compose up todo-app

run-dev:
	go run cmd/main.go

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database '${POSTGRES_URI}' up

migrate-down:
	migrate -path ./schema -database '${POSTGRES_URI}' down

swag:
	swag init -g cmd/main.go