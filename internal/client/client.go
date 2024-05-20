package client

import (
	"net/http"
	"time"

	"github.com/vinofsteel/pokedex-repl/internal/cache"
)

const baseURL string = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
	cache      cache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(cacheInterval),
	}
}
