package main

import (
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
)

const InternalServerError = `{"errorCode":"InternalServerError","errorMessage":"Internal Server Error"}`

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Start internalerror function")
	fmt.Println("Successfully finish internalerror function")
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: InternalServerError}, nil
}

func main() {
	basicLambda.Start(handler)
}
