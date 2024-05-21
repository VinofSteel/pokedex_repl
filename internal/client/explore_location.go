package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(param *string) (RespLocation, error) {
	if param == nil || *param == "" {
		return RespLocation{}, errors.New("you need to declare the name parameter to explore a location, like this: explore <name>")
	}
	url := baseURL + "location-area/" + *param

	if val, ok := c.cache.Get(url); ok {
		var location RespLocation

		if err := json.Unmarshal(val, &location); err != nil {
			return RespLocation{}, err
		}

		return location, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocation{}, err
	}

	location := RespLocation{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return RespLocation{}, fmt.Errorf("invalid location name, error: %v", err)
	}

	c.cache.Add(url, data)
	return location, nil
}
