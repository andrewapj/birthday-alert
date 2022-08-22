.PHONY: build

build:
	GOOS=linux GOARCH=amd64 go build -o main main.go
	zip main.zip main

test:
	docker-compose up -d
	go test ./... -v
	docker-compose down