package client

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		location := RespShallowLocations{}

		if err := json.Unmarshal(val, &location); err != nil {
			return RespShallowLocations{}, err
		}

		return location, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	location := RespShallowLocations{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)
	return location, nil
}
