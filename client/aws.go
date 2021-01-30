package client

import (
	"aws-cost-notify/model"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

type AwsClient struct {
	costexplorerSrv *costexplorer.CostExplorer
}

var _ AwsCostFetcher = (*AwsClient)(nil)

func NewAwsClient() *AwsClient {
	sess := session.Must(session.NewSession())
	costexplorerSrv := costexplorer.New(sess)

	return &AwsClient{
		costexplorerSrv: costexplorerSrv,
	}
}

const (
	dateFormatForSDK = "2006-01-02"
)

func (a *AwsClient) FetchCost(startDate, endDate time.Time, granularity Granularity) (*model.AwsCost, error) {
	metric := costexplorer.MetricNetUnblendedCost
	metrics := []*string{&metric}
	beforDateStr := startDate.Format(dateFormatForSDK)
	targetDateStr := endDate.Format(dateFormatForSDK)
	timePeriod := costexplorer.DateInterval{
        Start: &beforDateStr,
        End:   &targetDateStr,
    }

	inputParams := &costexplorer.GetCostAndUsageInput{
        Granularity: aws.String(granularity),
        Metrics:     metrics,
        TimePeriod:  &timePeriod,
    }
	ret, err := a.costexplorerSrv.GetCostAndUsage(inputParams)
	if err != nil {
		return nil, err
	}

	amount := ret.ResultsByTime[0].Total["NetUnblendedCost"].Amount

	return &model.AwsCost{
		StartDate: startDate,
		EndDate: endDate,
		Amount: *amount,
	}, nil
}
