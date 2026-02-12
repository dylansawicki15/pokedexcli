package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dylansawicki15/pokedexcli/pokecache"
)

var detailsCache = pokecache.NewCache(5 * time.Minute)

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocalizedName struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type EncounterMethodVersionDetail struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource               `json:"encounter_method"`
	VersionDetails  []EncounterMethodVersionDetail `json:"version_details"`
}

type EncounterDetail struct {
	Chance          int                `json:"chance"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	MaxLevel        int                `json:"max_level"`
	Method          NamedAPIResource   `json:"method"`
	MinLevel        int                `json:"min_level"`
}

type PokemonVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          NamedAPIResource  `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource       `json:"pokemon"`
	VersionDetails []PokemonVersionDetail `json:"version_details"`
}

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             NamedAPIResource      `json:"location"`
	Name                 string                `json:"name"`
	Names                []LocalizedName       `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

func fetchLocationAreaData(areaName string) ([]PokemonEncounter, error) {
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)

	if data, ok := detailsCache.Get(endpoint); ok {
		var locationAreaData LocationArea
		err := json.Unmarshal(data, &locationAreaData)
		if err != nil {
			fmt.Printf("Error unmarshalling cached data: %v\n", err)
			return nil, err
		}
		return locationAreaData.PokemonEncounters, nil
	}

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

	var locationAreaData LocationArea
	err = json.Unmarshal(data, &locationAreaData)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return nil, err
	}

	detailsCache.Add(endpoint, data)
	return locationAreaData.PokemonEncounters, nil
}
