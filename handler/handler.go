package handler

import (
	"aws-cost-notify/client"
	"aws-cost-notify/model"
	"aws-cost-notify/util/clock"
	"fmt"
	"log"
	"time"

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

	// 日別料金の取得
	oneDaysAgoCost, err := h.awsClient.FetchCost(today.AddDate(0, 0, -1), today, client.GranularityDaily)
	if err != nil {
		log.Print("failed to fetch cost err:", err)
		return InternalErrorResponse("failed to fetch cost error")
	}
	twoDaysAgoCost, err := h.awsClient.FetchCost(today.AddDate(0, 0, -2), today.AddDate(0, 0, -1), client.GranularityDaily)
	if err != nil {
		log.Print("failed to fetch cost err:", err)
		return InternalErrorResponse("failed to fetch cost error")
	}

	// 月間料金の取得。1日の場合は前月料金を取得。
	monthlyCost, err := func(today time.Time) (*model.AwsCost, error) {
		thisMonthFirst := clock.GetMonthFirst(today)
		if today.Day() == 1 {
			return h.awsClient.FetchCost(
				thisMonthFirst.AddDate(0, -1, 0),
				today,
				client.GranularityMonthy,
			)
		} else {
			return h.awsClient.FetchCost(
				thisMonthFirst,
				today,
				client.GranularityMonthy,
			)
		}
	}(today)
	if err != nil {
		log.Print("failed to fetch cost err:", err)
		return InternalErrorResponse("failed to fetch cost error")
	}

	fmt.Printf("oneDaysAgoCost: %+v\ntwoDaysAgoCost: %+v", oneDaysAgoCost, twoDaysAgoCost)

	outputMsg := fmt.Sprintf(`

### Daily Cost ###
%s: $%s
%s: $%s

### Monthly Cost ###
%s~%s: $%s
`,
	twoDaysAgoCost.StartDate.Format(outputDateFormat),
	twoDaysAgoCost.Amount,
	oneDaysAgoCost.StartDate.Format(outputDateFormat),
	oneDaysAgoCost.Amount,
	monthlyCost.StartDate.Format(outputDateFormat),
	monthlyCost.EndDate.AddDate(0, 0, -1).Format(outputDateFormat),
	monthlyCost.Amount,
	)

	fmt.Println(outputMsg)

	if err := h.noticeClient.Notify(outputMsg); err != nil {
		log.Print("failed to notify err:", err)
		return InternalErrorResponse("notify error")
	}

	return SuccessResponse()
}
