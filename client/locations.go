package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}

func GetLocations() {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%slocation-area", BaseURL), nil)
	if err != nil {
		fmt.Printf("Error generating a new GetLocations request: %v \n", err)
		return
	}

	res, err := Client.Do(req)
	if err != nil {
		fmt.Printf("Error getting locations data from PokeAPI: %v \n", err)
		return
	}
	fmt.Println(res, "REQUISITION RESPONSE TO POKE API")
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("Unexpected status code:", res.StatusCode)
	}

	var location Location
	if err := json.NewDecoder(res.Body).Decode(&location); err != nil {
		fmt.Printf("Error decoding location response body from PokeAPI: %v \n", err)
	}

	fmt.Println(location, "LOCATION")
}
