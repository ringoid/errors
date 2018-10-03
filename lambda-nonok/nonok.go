package main

import (
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Start nonok function")
	fmt.Println("Successfully finish nonok function")
	return events.APIGatewayProxyResponse{StatusCode: 503}, nil
}

func main() {
	basicLambda.Start(handler)
}
