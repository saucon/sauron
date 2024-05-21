package notify_error

import (
	"github.com/go-resty/resty/v2"
	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"time"
)

type Client struct {
	client *resty.Client
	config *logconfig.GspaceChat
}

var url = "https://chat.googleapis.com/v1/spaces"

func New(config *logconfig.GspaceChat) *Client {
	return &Client{
		client: resty.New().
			SetBaseURL(url).
			SetTimeout(1 * time.Minute).
			SetRetryCount(3).
			SetRetryMaxWaitTime(5 * time.Second).
			SetRetryWaitTime(5 * time.Second),
		config: config,
	}
}
