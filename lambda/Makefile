.PHONY: build

build:
	GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bootstrap main.go
	zip bootstrap.zip bootstrap

test:
	docker-compose up -d
	go test ./... -v
	docker-compose down
