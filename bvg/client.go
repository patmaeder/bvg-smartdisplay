package bvg

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: "https://v6.bvg.transport.rest",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
