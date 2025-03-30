# Load .env variables
include .env

install:
	go mod tidy

build:
	go build cmd/app/main.go

run: 
	go run cmd/app/main.go

swagger:
	swag init -g cmd/app/main.go -o cmd/app/docs

up:
	migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -path migrations up

down:
	migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -path migrations down

start:
	make up && make swagger && make run