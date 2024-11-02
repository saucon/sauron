package notify_error

import (
	"fmt"
	"net/http"
)

func (c *Client) SendNotif(req NotifyRequest) (err error) {
	r, err := c.client.R().
		SetBody(req).
		SetDebug(true).
		Post(fmt.Sprintf("/%s/messages?key=%s&token=%s", c.config.SpaceID, c.config.SpaceSecret, c.config.SpaceToken))
	if err != nil {
		return
	}

	if r.StatusCode() != http.StatusOK {
		return
	}

	return
}
