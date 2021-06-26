package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda/messages"
	lc "github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

type Input struct {
	Port          int
	Payload       interface{}
	ClientContext *lc.ClientContext
	Deadline      *messages.InvokeRequest_Timestamp
}

type Event struct {
	TraceId       string `json:"traceId"`
	LogGroupName  string `json:"logGroupName"`
	LogStreamName string `json:"logStreamName"`
}

type Response struct {
	TraceId string `json:"traceId"`
	Message string `json:"message"`
}

func main() {
	event := Event{
		TraceId:       "12345",
		LogGroupName:  "/aws/lambda/ecdr-dev-commonwell-token-generator-lambda",
		LogStreamName: "2021/06/26/[$LATEST]c0849761e67e43b195e4f66eda8d9e2f",
	}

	resp, err := golambdainvoke.Run(Input{
		Port:    8080,
		Payload: event,
	})

	if err != nil {
		fmt.Println("Failed to invoke the lambda: ", err)
	}

	var response Response
	err = json.Unmarshal(resp, &response)
	fmt.Println("Trace ID: ", response.TraceId)
}
