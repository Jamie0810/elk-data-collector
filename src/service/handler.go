package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type Event struct {
	TraceId       string `json:"traceId"`
	LogGroupName  string `json:"logGroupName"`
	LogStreamName string `json:"logStreamName"`
}

type Response struct {
	TraceId string `json:"traceId"`
	Message string `json:"message"`
}

func LambdaHandler(ctx context.Context, event Event) (Response, error) {
	fmt.Println("-----Lambda execution started-----")

	var limit int64 = 100

	logGroupName := event.LogGroupName
	logStreamName := event.LogStreamName

	if logGroupName == "" || logStreamName == "" {
		fmt.Println("You must supply a log group name (-g LOG-GROUP) and log stream name (-s LOG-STREAM)")
		return Response{Message: "Job failed"}, nil
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	resp, err := GetLogEvents(sess, limit, logGroupName, logStreamName)
	if err != nil {
		fmt.Println("Got error getting log events: ", err)
		return Response{Message: "Job failed"}, err
	}

	fmt.Println("Event messages for stream " + logStreamName + " in log group  " + logGroupName)

	for _, event := range resp.Events {
		fmt.Println("------Message-------")
		fmt.Println(*event.Message)
	}

	return Response{
		TraceId: event.TraceId,
		Message: "Job completed",
	}, nil
}

func GetLogEvents(sess *session.Session, limit int64, logGroupName string, logStreamName string) (*cloudwatchlogs.GetLogEventsOutput, error) {
	svc := cloudwatchlogs.New(sess)

	resp, err := svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		Limit:         &limit,
		LogGroupName:  &logGroupName,
		LogStreamName: &logStreamName,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
