package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Hello my World!",
		}, nil
	})
}
