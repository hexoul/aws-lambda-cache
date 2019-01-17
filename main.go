package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	cmc "github.com/hexoul/go-coinmarketcap"
	"github.com/hexoul/go-coinmarketcap/types"
)

const (
	cmcAPIKey    = "CMC_API_KEY"
	targetSymbol = "BTC"
	targetQuotes = "USD,KRW"
)

var (
	headers = map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Methods": "*",
	}
	cacheMap map[string]string

	lastUpdatedPrice time.Time
	lastPriceInfo    string
)

func init() {
	cacheMap = make(map[string]string)
	lastUpdatedPrice = time.Now().AddDate(0, 0, -1)
}

func cache(request events.APIGatewayProxyRequest) (respBody string) {
	key := request.QueryStringParameters["key"]
	if request.Body != "" {
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
	return
}

func price() (respBody string) {
	if time.Now().Sub(lastUpdatedPrice).Minutes() < 1 {
		return lastPriceInfo
	}
	if data, err := cmc.GetInstanceWithKey(cmcAPIKey).CryptoMarketQuotesLatest(&types.Options{
		Symbol:  targetSymbol,
		Convert: targetQuotes,
	}); err == nil {
		for _, v := range data.CryptoMarket {
			if last, err := time.Parse(time.RFC3339, v.Quote["USD"].LastUpdated); err == nil {
				lastUpdatedPrice = last
			}
			if b, bErr := json.Marshal(v); bErr == nil {
				lastPriceInfo = string(b)
			}
			break
		}
	}
	return lastPriceInfo
}

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	respBody := ""
	statusCode := 200

	if request.QueryStringParameters["price"] != "" {
		respBody = price()
	} else if request.QueryStringParameters["key"] != "" {
		respBody = cache(request)
	}

	return events.APIGatewayProxyResponse{Headers: headers, Body: respBody, StatusCode: statusCode}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
