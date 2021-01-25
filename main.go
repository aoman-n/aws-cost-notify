package main

import (
	"aws-billing-notify/client"
	"aws-billing-notify/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	awsClient := client.NewAwsClient()
	lineClient := client.NewLineClient()
	handler := handler.New(awsClient, lineClient)
	lambda.Start(handler.Run)
}