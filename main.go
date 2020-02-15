package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func() (events.APIGatewayProxyResponse, error) {
		log.Println("Processing Lambda request!")

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Hello your World!",
		}, nil
	})
}
