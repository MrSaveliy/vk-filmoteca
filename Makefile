APP_NAME=vk-filmoteca

DB_CONNECTION_STRING=user=postgres password=postgres dbname=postgres sslmode=disable

migrate-up:
    goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up

migrate-down:
    goose postgres "$(DB_CONNECTION_STRING)" down

build:
    go build -o $(APP_NAME) main.go

run:
    go run main.go

test:
    go test ./...

.PHONY: migrate build run test