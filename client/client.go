package client

import (
	"aws-cost-notify/model"
	"time"

	"github.com/aws/aws-sdk-go/service/costexplorer"
)

type Granularity = string

const (
	GranularityDaily Granularity = costexplorer.GranularityDaily
	GranularityMonthy Granularity = costexplorer.GranularityMonthly
)


type AwsCostFetcher interface {
	FetchCost(startDate, endDate time.Time, granularity Granularity) (*model.AwsCost, error)
}

type Notifier interface {
	Notify(msg string) error
}
