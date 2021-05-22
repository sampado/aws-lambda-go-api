package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sampado/aws-lambda-go-api/app/hello"
)

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       hello.HelloMessage("Basic Lambda"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
