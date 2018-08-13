package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestCache(t *testing.T) {
	event := events.APIGatewayProxyRequest{}
	event.QueryStringParameters = make(map[string]string)
	event.QueryStringParameters["key"] = "test"
	event.QueryStringParameters["val"] = "123"
	resp, err := lambdaHandler(nil, event)
	if err != nil {
		t.Errorf("Failed to set")
	}

	event = events.APIGatewayProxyRequest{}
	event.QueryStringParameters = make(map[string]string)
	event.QueryStringParameters["key"] = "test"
	resp, err = lambdaHandler(nil, event)
	if resp.Body == "" || err != nil {
		t.Errorf("Failed to get")
	}
	t.Log(resp.Body)
}
