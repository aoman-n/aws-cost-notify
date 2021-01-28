package client

import "fmt"


type LineClient struct {
	lineNotifyToken string
	linePostUrl string
}

var _ Notifier = (*LineClient)(nil)

func NewLineClient(lineNotifyToken, linePostUrl string) *LineClient {
	return &LineClient{
		lineNotifyToken: lineNotifyToken,
		linePostUrl: linePostUrl,
	}
}

func (l *LineClient) Notify(msg string) error {
	fmt.Println(msg)
	return nil
}
