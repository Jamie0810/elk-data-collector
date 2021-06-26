package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	TraceId string `json:"traceId"`
}

type Response struct {
	TraceId string `json:"traceId"`
	Message string `json:"message"`
}

func LambdaHandler(ctx context.Context, event Event) (Response, error) {
	return Response{
		TraceId: event.TraceId,
		Message: "Job completed",
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
