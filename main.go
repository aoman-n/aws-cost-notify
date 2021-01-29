package main

import (
	"aws-cost-notify/client"
	"aws-cost-notify/handler"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	LINE_NOTIFY_TOKEN := os.Getenv("LINE_NOTIFY_TOKEN")
	LINE_POST_URL := os.Getenv("LINE_POST_URL")
	if LINE_NOTIFY_TOKEN == "" {
		log.Fatalf(`required env "LINE_NOTIFY_TOKEN"`)
	}
	if LINE_POST_URL == "" {
		log.Fatalf(`required env "LINE_POST_URL"`)
	}

	awsClient := client.NewAwsClient()
	lineClient := client.NewLineClient(LINE_NOTIFY_TOKEN, LINE_POST_URL)
	handler := handler.New(awsClient, lineClient)
	lambda.Start(handler.Run)
}