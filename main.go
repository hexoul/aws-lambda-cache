package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	cacheMap map[string]string
)

func init() {
	cacheMap = make(map[string]string)
}

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	respBody := ""
	statusCode := 200

	if request.QueryStringParameters != nil {
		key := request.QueryStringParameters["key"]
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

	return events.APIGatewayProxyResponse{Body: respBody, StatusCode: statusCode}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
