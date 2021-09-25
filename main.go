package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func HandleRequest() error {
	log.Println("Lambda started")

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
