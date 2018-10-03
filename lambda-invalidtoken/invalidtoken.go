package main

import (
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
)

const InvalidAccessTokenClientError = `{"errorCode":"InvalidAccessTokenClientError","errorMessage":"Invalid access token"}`

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Start invalidtoken function")
	fmt.Println("Successfully finish invalidtoken function")
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: InvalidAccessTokenClientError}, nil
}

func main() {
	basicLambda.Start(handler)
}
