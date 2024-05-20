package client

import (
	"net/http"
	"time"
)

const BaseURL string = "https://pokeapi.co/api/v2/"

var Client *http.Client = &http.Client{
	Timeout: 20 * time.Second,
}