.PHONY: run build

run:
	go run cmd/api/main.go

build:
	go build -o bin/app cmd/api/main.go

dev: run