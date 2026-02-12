package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var catchRNG = rand.New(rand.NewSource(time.Now().UnixNano()))

type Pokemon struct {
	Data pokemonDataResponse
}

var pokedex = map[string]Pokemon{}

type pokemonDataResponse struct {
	Abilities              json.RawMessage `json:"abilities"`
	BaseExperience         int             `json:"base_experience"`
	Cries                  json.RawMessage `json:"cries"`
	Forms                  json.RawMessage `json:"forms"`
	GameIndices            json.RawMessage `json:"game_indices"`
	Height                 int             `json:"height"`
	HeldItems              json.RawMessage `json:"held_items"`
	ID                     int             `json:"id"`
	IsDefault              bool            `json:"is_default"`
	LocationAreaEncounters string          `json:"location_area_encounters"`
	Moves                  json.RawMessage `json:"moves"`
	Name                   string          `json:"name"`
	Order                  int             `json:"order"`
	PastAbilities          json.RawMessage `json:"past_abilities"`
	PastStats              json.RawMessage `json:"past_stats"`
	PastTypes              json.RawMessage `json:"past_types"`
	Species                json.RawMessage `json:"species"`
	Sprites                json.RawMessage `json:"sprites"`
	Stats                  json.RawMessage `json:"stats"`
	Types                  json.RawMessage `json:"types"`
	Weight                 int             `json:"weight"`
}

func commandCatch(name string) error {
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return fmt.Errorf("pokemon name cannot be empty")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemonData, err := fetchPokemonData(name)
	if err != nil {
		return err
	}

	catchChance := calculateCatchChance(pokemonData.BaseExperience)
	roll := catchRNG.Intn(100) + 1

	if roll <= catchChance {
		fmt.Printf("%s was caught!\n", name)
		pokedex[name] = Pokemon{
			Data: pokemonData,
		}
		return nil
	}

	fmt.Printf("%s escaped!\n", name)
	return nil
}

func calculateCatchChance(baseExperience int) int {
	chance := 100 - (baseExperience / 2)
	if chance < 5 {
		return 5
	}
	if chance > 95 {
		return 95
	}
	return chance
}

func fetchPokemonData(name string) (pokemonDataResponse, error) {
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", name)

	if data, ok := cache.Get(endpoint); ok {
		var pokemonData pokemonDataResponse
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			return pokemonDataResponse{}, err
		}
		return pokemonData, nil
	}

	res, err := http.Get(endpoint)
	if err != nil {
		return pokemonDataResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return pokemonDataResponse{}, fmt.Errorf("unexpected status %d: %s", res.StatusCode, string(body))
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonDataResponse{}, err
	}

	var pokemonData pokemonDataResponse
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return pokemonDataResponse{}, err
	}

	cache.Add(endpoint, data)
	return pokemonData, nil
}
