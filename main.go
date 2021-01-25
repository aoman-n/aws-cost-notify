package main

import (
	"aws-billing-notify/clients"
	"aws-billing-notify/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	awsClient := clients.NewAwsClient()
	lineClient := clients.NewLineClient()
	handler := handler.New(awsClient, lineClient)
	lambda.Start(handler.Run)
}