package handler

import "github.com/aws/aws-lambda-go/events"

func successResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func internalErrorResponse(msg string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body: msg,
	}, nil
}
