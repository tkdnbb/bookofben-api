.PHONY: run build

run:
	go run cmd/api/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/app cmd/api/main.go

buildfc:
	GOOS=linux GOARCH=amd64 go build -o bin/main cmd/fc/main.go

dev: run