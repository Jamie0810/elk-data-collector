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
	TraceId string `json:"traceId"`
}

type Response struct {
	TraceId string `json:"traceId"`
	Message string `json:"message"`
}

func main() {
	handlerEvent := Event{TraceId: "12345"}

	resp, err := golambdainvoke.Run(Input{
		Port:    8001,
		Payload: handlerEvent,
	})

	if err != nil {
		fmt.Println("Failed to invoke the lambda: ", err)
	}

	var response Response
	err = json.Unmarshal(resp, &response)
	fmt.Println("Trace ID: ", response.TraceId)
}
