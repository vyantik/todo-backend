include .env
export

build:
	go build -o todo-app cmd/main.go

run:
	./todo-app

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