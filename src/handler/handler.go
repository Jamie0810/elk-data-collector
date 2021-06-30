package handler

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type AwsLogs struct {
	Data string `json:"data"`
}
type Event struct {
	AwsLogs AwsLogs `json:"awslogs"`
}
type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type LogEvent struct {
	ID        string `json:"id"`
	Timestamp int    `json:"timestamp"`
	Message   string `json:"message"`
}
type Message struct {
	MessageType         string     `json:"messageType"`
	Owner               string     `json:"owner"`
	LogGroup            string     `json:"logGroup"`
	LogStream           string     `json:"logStream"`
	SubscriptionFilters []string   `json:"subscriptionFilters"`
	LogEvents           []LogEvent `json:"logEvents"`
}

func LambdaHandler(ctx context.Context, event Event) (Response, error) {
	fmt.Println("-----Lambda execution started-----")

	// Get compressed data
	compressedData := event.AwsLogs.Data
	data, err := base64.StdEncoding.DecodeString(compressedData)
	if err != nil {
		return Response{Message: "Job Failed", Error: err.Error()}, nil
	}
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return Response{Message: "Job Failed", Error: err.Error()}, nil
	}
	defer zr.Close()

	dec := json.NewDecoder(zr)
	m := Message{}
	err = dec.Decode(&m)
	fmt.Println(m.LogEvents[0].Message)

	return Response{
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
