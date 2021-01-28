package handler

import (
	"aws-billing-notify/client"
	"aws-billing-notify/util/clock"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	awsClient client.AwsCostFetcher
	noticeClient client.Notifier
}

func New(awsClient client.AwsCostFetcher, noticeClient client.Notifier) *Handler {
	return &Handler{
		awsClient: awsClient,
		noticeClient: noticeClient,
	}
}

const outputDateFormat = "2006/01/02"

func (h *Handler) Run(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	today := clock.JstNow()

	oneDaysAgoCost, err := h.awsClient.FetchCost(today.AddDate(0, 0, -1), today)
	if err != nil {
		return InternalErrorResponse("fetch cost error")
	}
	twoDaysAgoCost, err := h.awsClient.FetchCost(today.AddDate(0, 0, -2), today.AddDate(0, 0, -1))
	if err != nil {
		return InternalErrorResponse("fetch cost error")
	}

	fmt.Printf("oneDaysAgoCost: %+v\ntwoDaysAgoCost: %+v", oneDaysAgoCost, twoDaysAgoCost)

	outputMsg := fmt.Sprintf(`## AWS利用料金 ##

### Daily Cost ###
%s: $%s
%s: $%s

### Monthly Cost ###
`,
	twoDaysAgoCost.StartDate.Format(outputDateFormat),
	twoDaysAgoCost.Amount,
	oneDaysAgoCost.StartDate.Format(outputDateFormat),
	oneDaysAgoCost.Amount,
	)

	if err := h.noticeClient.Notify(outputMsg); err != nil {
		log.Print("failed to notify err:", err)
		return InternalErrorResponse("notify error")
	}

	return SuccessResponse()
}

// Line Output Message
// ## AWS利用料金 ##
//
// ### Daily Cost ###
// 2020-01-18: $2.10
// 2020-01-19: $1.10
//
// ### Monthly Cost ###
// 2020-01-01~2020-01-20: $40.10
