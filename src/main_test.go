package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaHandler(t *testing.T) {
	ctx := context.Background()

	event := Event{
		TraceId: "123",
	}

	resp, err := LambdaHandler(ctx, event)
	if err != nil {
		t.Fatalf("Job failed: %v", err)
	}
	assert.Equal(t, event.TraceId, resp.TraceId)
}
