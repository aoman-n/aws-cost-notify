package handler

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

func Run(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("invoked handler")

	LINE_TOKEN := os.Getenv("LINE_NOTIFY_TOKEN")
	LINE_POST_URL := os.Getenv("LINE_POST_URL")
	fmt.Printf("LINE_TOKEN: %s \nLINE_POST_URL: %s\n", LINE_TOKEN, LINE_POST_URL)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
