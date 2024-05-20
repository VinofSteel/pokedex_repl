package client

import (
	"net/http"
	"time"
)

const baseURL string = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
