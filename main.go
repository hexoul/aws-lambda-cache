package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	cacheMap map[string]string
	headers  = map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": "*",
	}
)

func init() {
	cacheMap = make(map[string]string)
}

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	respBody := ""
	statusCode := 200

	key := request.QueryStringParameters["key"]
	if key == "" {
		// Nothing to do
	} else if request.Body != "" {
		// In case of "set" with body
		cacheMap[key] = request.Body
	} else if request.QueryStringParameters != nil {
		val := request.QueryStringParameters["val"]
		if val == "" {
			// In case of "get"
			respBody = cacheMap[key]
			cacheMap[key] = ""
		} else if key != "" {
			// In case of "set"
			cacheMap[key] = val
		}
	}

	return events.APIGatewayProxyResponse{Headers: headers, Body: respBody, StatusCode: statusCode}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
