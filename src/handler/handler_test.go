package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaHandler(t *testing.T) {
	ctx := context.Background()

	event := Event{
		TraceId:       "12345",
		LogGroupName:  "/aws/lambda/ecdr-dev-commonwell-token-generator-lambda",
		LogStreamName: "2021/06/26/[$LATEST]c0849761e67e43b195e4f66eda8d9e2f",
	}

	resp, err := LambdaHandler(ctx, event)
	if err != nil {
		t.Fatalf("Job failed: %v", err)
	}
	assert.Equal(t, "Job completed", resp.Message)
}
