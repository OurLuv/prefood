.SILENT:

run:
	go run cmd/main.go

build:
	docker-compose up -d --build 