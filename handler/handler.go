package handler

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Run(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("invoked handler")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
