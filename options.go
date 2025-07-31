package metrica

import (
	"net/http"
	"time"
)

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithAuthToken(authToken string) Option {
	return func(c *Client) {
		c.authToken = authToken
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout * time.Second
	}
}
