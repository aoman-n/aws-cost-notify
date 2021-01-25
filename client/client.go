package client

import "aws-billing-notify/domain"

type AwsBillinger interface {
	FetchBilling() (*domain.Billing, error)
}

type Notifier interface {
	Notify(*domain.Billing) error
}
