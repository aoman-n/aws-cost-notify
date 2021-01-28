package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)


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
	postData := url.Values{}
	postData.Add("message", msg)
	req, err := http.NewRequest("POST", l.linePostUrl, strings.NewReader(postData.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", l.lineNotifyToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("line response: ", resp)

	return nil
}
