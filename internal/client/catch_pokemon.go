package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(param *string) (Pokemon, error) {
	if param == nil || *param == "" {
		return Pokemon{}, errors.New("you need to declare the name parameter to catch a pokémon, like this: catch <name>")
	}
	url := baseURL + "pokemon/" + *param

	if val, ok := c.cache.Get(url); ok {
		var pokemon Pokemon

		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("invalid pokemon name, error: %v", err)
	}

	c.cache.Add(url, data)
	return pokemon, nil
}
