package client

import (
	"aws-billing-notify/domain"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type AwsClient struct {
	cloudwatchSrv *cloudwatch.CloudWatch
}

var _ AwsBillinger = (*AwsClient)(nil)

func NewAwsClient() *AwsClient {
	sess := session.Must(session.NewSession())
	srv := cloudwatch.New(
		sess,
		aws.NewConfig().WithRegion("us-east-1"),
	)
	return &AwsClient{
		cloudwatchSrv: srv,
	}
}

func (a *AwsClient) FetchBilling() (*domain.Billing, error) {
	params := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []*cloudwatch.Dimension{
            {
                Name:  aws.String("Currency"),
                Value: aws.String("USD"),
            },
        },
        StartTime:  aws.Time(time.Now().Add(time.Hour * -24)),
        EndTime:    aws.Time(time.Now()),
        Period:     aws.Int64(86400),
        Namespace:  aws.String("AWS/Billing"),
        MetricName: aws.String("EstimatedCharges"),
        Statistics: []*string{
            aws.String(cloudwatch.StatisticMaximum),
        },
    }
    resp, err := a.cloudwatchSrv.GetMetricStatistics(params)
    if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(strings.Repeat("*", 100))
	fmt.Printf("resp: %+v\n", resp)
	fmt.Println(strings.Repeat("*", 100))
	mockBilling := &domain.Billing{
		CurrentMonthTotal: "hoge",
		PreviousDay: "hoge",
	}
	return mockBilling, nil
}
