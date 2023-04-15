.SILENT:

default: build

build:
	go build -race -o .bin/app.go cmd/app.go

stop:
	docker compose down -v

run: stop build
	docker compose up --build
