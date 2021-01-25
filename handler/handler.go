package handler

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

func Run(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("invoked handler")

	key := os.Getenv("YOUTUBE_API_KEY")
	fmt.Printf("key: %s \n", key)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
