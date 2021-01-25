package handler

import (
	"aws-billing-notify/client"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	awsClient client.AwsBillinger
	noticeClient client.Notifier
}

func New(awsClient client.AwsBillinger, noticeClient client.Notifier) *Handler {
	return &Handler{
		awsClient: awsClient,
		noticeClient: noticeClient,
	}
}

func (h *Handler) Run(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	LINE_TOKEN := os.Getenv("LINE_NOTIFY_TOKEN")
	LINE_POST_URL := os.Getenv("LINE_POST_URL")
	fmt.Printf("LINE_TOKEN: %s \nLINE_POST_URL: %s\n", LINE_TOKEN, LINE_POST_URL)

	billing, err := h.awsClient.FetchBilling()
	if err != nil {
		log.Print("failed to fetch billing err:", err)
		return InternalErrorResponse("fetch billing error")
	}

	if err := h.noticeClient.Notify(billing); err != nil {
		log.Print("failed to notify err:", err)
		return InternalErrorResponse("fetch billing error")
	}

	return SuccessResponse()
}
