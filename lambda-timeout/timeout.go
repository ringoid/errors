package main

import (
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"time"
	"fmt"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Start timeout function")
	time.Sleep(4 * time.Minute)
	fmt.Println("Successfully finish timeout function")
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	basicLambda.Start(handler)
}
