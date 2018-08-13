package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestCache(t *testing.T) {
	event := events.APIGatewayProxyRequest{}
	event.QueryStringParameters = make(map[string]string)
	event.QueryStringParameters["key"] = "0x25c317c8d0a63c122073ae52984e8477e7fbc322c93a9457c5579ee6e5a813b3"
	resp, err := lambdaHandler(nil, event)
	if resp.Body == "" || err != nil {
		t.Errorf("Failed to start main")
	}
	t.Log(resp.Body)
}
