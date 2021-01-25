package client

import "aws-billing-notify/domain"

type LineClient struct {}

var _ Notifier = (*LineClient)(nil)

func NewLineClient() *LineClient {
	return &LineClient{}
}

func (l *LineClient) Notify(billing *domain.Billing) error {
	return nil
}
