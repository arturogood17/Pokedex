package pokeapi

import (
	"net/http"
	"time"
)

type Client struct { // http.Client ya tiene timeout
	HttpClient http.Client
}

func NewClient(timeout time.Duration) *Client {
	c := &Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
	return c
}
