package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func commandInspect(name string) error {
	p, exists := pokedex[name]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}

	pokemonData := p.Data

	type statEntry struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	}

	type typeEntry struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	}

	var stats []statEntry
	if len(pokemonData.Stats) > 0 {
		if err := json.Unmarshal(pokemonData.Stats, &stats); err != nil {
			return err
		}
	}

	var types []typeEntry
	if len(pokemonData.Types) > 0 {
		if err := json.Unmarshal(pokemonData.Types, &types); err != nil {
			return err
		}
	}

	order := []string{"hp", "attack", "defense", "special-attack", "special-defense", "speed"}

	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)

	fmt.Printf("Stats:\n")
	for _, name := range order {
		val := 0
		for _, s := range stats {
			if s.Stat.Name == name {
				val = s.BaseStat
				break
			}
		}
		fmt.Printf("  -%s: %d\n", name, val)
	}

	sort.Slice(types, func(i, j int) bool { return types[i].Slot < types[j].Slot })
	fmt.Printf("Types:\n")
	for _, t := range types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
