package client

import (
	"net/http"
	"time"
)

type Client struct {
	HttpClient http.Client
}

func New(timeout time.Duration) *Client {
	return &Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}
