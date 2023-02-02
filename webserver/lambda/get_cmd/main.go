package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// naked example of isolated entity handling via AWS lambda
// https://github.com/cubesatplatform/webserver/blob/main/webserver/myweb/myweb.go#L296

// https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
//

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cubesatplatform/webserver/webserver/mydb"
)

func main() {
	lambda.Start(handleHttpRequest)
}

func handleHttpRequest(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !strings.HasPrefix(request.Path, "/getcmd") {
		return events.APIGatewayProxyResponse{
			IsBase64Encoded: false,
			StatusCode:      http.StatusBadRequest,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            "invalid handler for this endpoint",
		}, nil
	}

	switch request.HTTPMethod {
	// skip OPTION - this is browser pre fly call
	case http.MethodOptions:
		return events.APIGatewayProxyResponse{
			IsBase64Encoded: false,
			StatusCode:      http.StatusOK,
			Headers:         map[string]string{"Content-Type": "application/json"},
			Body:            "",
		}, nil
	// proceed business logic here
	case http.MethodGet:

		bsid := request.QueryStringParameters["bsid"]
		if bsid == "" {
			return events.APIGatewayProxyResponse{
				IsBase64Encoded: false,
				StatusCode:      http.StatusBadRequest,
				Headers:         map[string]string{"Content-Type": "text/plain"},
				Body:            "bsid query param missed",
			}, nil
		}

		data := mydb.GetCMD("bsid")
		jsonByteArray, err := json.Marshal(data)

		if err != nil {
			return events.APIGatewayProxyResponse{
				IsBase64Encoded: false,
				StatusCode:      http.StatusBadRequest,
				Headers:         map[string]string{"Content-Type": "text/plain"},
				Body:            err.Error(),
			}, nil
		}

		// return result as JSON
		return events.APIGatewayProxyResponse{
			IsBase64Encoded: false,
			StatusCode:      http.StatusOK,
			Headers:         map[string]string{"Content-Type": "application/json"},
			Body:            string(jsonByteArray),
		}, nil

	default:
		return events.APIGatewayProxyResponse{
			IsBase64Encoded: false,
			StatusCode:      http.StatusBadRequest,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            "request method not supported",
		}, nil
	}
}
