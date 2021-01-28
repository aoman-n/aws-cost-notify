package client

import (
	"aws-billing-notify/model"
	"time"
)


type AwsCostFetcher interface {
	FetchCost(startDate, endDate time.Time) (*model.AwsCost, error)
}

type Notifier interface {
	Notify(msg string) error
}
