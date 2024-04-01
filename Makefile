include .env
export

build:
	@go build -o bin/api cmd/server.go

run:
	@./bin/api

test:
	@go test ./..

tidy:
	@go mod tidy

migrate:
	@migrate -path db/migrations -database "$(DB_URL)" -verbose up

drop:
	@migrate -path db/migrations -database "$(DB_URL)" -verbose down

create:
	@migrate create -ext sql -dir db/migrations rentio_tables

up:
	docker-compose up --build

down:
	docker-compose down