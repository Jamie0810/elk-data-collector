package main

import (
	"data-collector/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.LambdaHandler)
}
