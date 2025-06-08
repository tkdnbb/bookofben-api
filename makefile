.PHONY: run build

run:
	go run cmd/api/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/app cmd/api/main.go

dev: run