package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	HttpClient http.Client
	Timeout    time.Duration
}

func NewClient(timeout time.Duration) *Client {
	c := &Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
	return c
}
