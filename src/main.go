package main

import (
	service "data-collector/service"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(service.LambdaHandler)
}
