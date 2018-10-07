package main

import (
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
)

var TooOldAppVersionClientError = `{"errorCode":"TooOldAppVersionClientError","errorMessage":"Too old app version"}`

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Start old function")
	fmt.Println("Successfully finihs old function")
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: TooOldAppVersionClientError}, nil
}

func main() {
	basicLambda.Start(handler)
}
