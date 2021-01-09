package main

import (
	"context"

	"github.com/Uchencho/serverlessTest/internal"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
)

type awsEventHandlerFunc func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func getAWSLambdaEventHandler() awsEventHandlerFunc {
	a := internal.New()
	lambdaProxyAdapter := handlerfunc.New(a.Handler())
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return lambdaProxyAdapter.ProxyWithContext(ctx, req)
	}
}

func main() {
	h := getAWSLambdaEventHandler()
	lambda.Start(h)

}
