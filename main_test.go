package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"

	cmc "github.com/hexoul/go-coinmarketcap"
	"github.com/hexoul/go-coinmarketcap/types"
)

func TestCache(t *testing.T) {
	event := events.APIGatewayProxyRequest{}
	event.QueryStringParameters = make(map[string]string)
	event.QueryStringParameters["key"] = "test"
	event.QueryStringParameters["val"] = "123"
	resp, err := lambdaHandler(nil, event)
	if err != nil {
		t.Error("Failed to set")
	}

	event = events.APIGatewayProxyRequest{}
	event.QueryStringParameters = make(map[string]string)
	event.QueryStringParameters["key"] = "test"
	resp, err = lambdaHandler(nil, event)
	if resp.Body == "" || err != nil {
		t.Error("Failed to get")
	}
	t.Log(resp.Body)
}

func TestPrice(t *testing.T) {
	event := events.APIGatewayProxyRequest{}
	event.QueryStringParameters = make(map[string]string)
	event.QueryStringParameters["price"] = "BTC"
	resp, err := lambdaHandler(nil, event)
	if err != nil {
		t.Error("Failed to get price")
	}
	t.Log(resp.Body)

	lambdaHandler(nil, event)
	lambdaHandler(nil, event)
}

func TestCmcAPI(t *testing.T) {
	if data, err := cmc.GetInstanceWithKey("CMC_API_KEY").CryptoMarketQuotesLatest(&types.Options{
		Symbol:  "BTC",
		Convert: "USD,ETH",
	}); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range data.CryptoMarket {
			if b, bErr := json.Marshal(v); bErr == nil {
				t.Log(string(b))
			}
			if last, err := time.Parse(time.RFC3339, v.Quote["USD"].LastUpdated); err == nil {
				t.Logf("last: %s\n", last.String())
				t.Logf("diff: %f\n", time.Now().Sub(last).Minutes())
			}
		}
	}
}
