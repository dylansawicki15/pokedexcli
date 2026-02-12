package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var index = 1
var pageSize = 20

type locationAreaResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []locationAreaItem `json:"results"`
}

type locationAreaItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func fetchPokemonLocationAreas(back bool) ([]locationAreaItem, error) {
	if back {
		if index > 1 {
			index--
		}
	} else {
		index++
	}

	offset := (index - 2) * pageSize
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=%d", offset, pageSize)

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error fetching location area: %v\n", err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", res.StatusCode, string(body))
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	var locationAreaData locationAreaResponse
	err = json.Unmarshal(data, &locationAreaData)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return nil, err
	}

	return locationAreaData.Results, nil
}
