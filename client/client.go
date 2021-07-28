package client

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type (
	Client struct {
		retryClient *retryablehttp.Client
	}
)

// New will return a pointer to a new initialized service client.
func New(options ...func(*Client)) *Client {
	svc := &Client{
		retryClient: retryablehttp.NewClient(),
	}

	// disable default logger
	svc.retryClient.Logger = nil

	for _, option := range options {
		option(svc)
	}

	return svc
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	r, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, err
	}

	return c.retryClient.Do(r)
}

func WithRetryMax(retries int) func(*Client) {
	return func(c *Client) {
		c.retryClient.RetryMax = retries
	}
}

func WithRetryWaitMax(d time.Duration) func(*Client) {
	return func(c *Client) {
		c.retryClient.RetryWaitMax = d
	}
}

func WithRetryWaitMin(d time.Duration) func(*Client) {
	return func(c *Client) {
		c.retryClient.RetryWaitMin = d
	}
}

func WithLogger(l interface{}) func(*Client) {
	return func(c *Client) {
		c.retryClient.Logger = l
	}
}
