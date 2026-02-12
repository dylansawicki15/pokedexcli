package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var build = "0"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex (build %v)> ", build)
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Please enter a command.")
			continue
		}
		switch input[0] {
		case "exit":
			supportedCommands["exit"].callback()
		case "help":
			supportedCommands["help"].callback()
		case "map":
			supportedCommands["map"].callback()
		case "mapb":
			supportedCommands["mapb"].callback()
		case "explore":
			if len(input) < 2 {
				fmt.Println("Please provide a location area to explore.")
				continue
			}
			areaName := input[1]
			err := commandExplore(areaName)
			if err != nil {
				fmt.Printf("Error exploring area: %v\n", err)
			}
		case "catch":
			if len(input) < 2 {
				fmt.Println("Please provide a pokemon name to catch.")
				continue
			}
			pokemonName := input[1]
			err := commandCatch(pokemonName)
			if err != nil {
				fmt.Printf("Error catching pokemon: %v\n", err)
			}
		case "inspect":
			if len(input) < 2 {
				fmt.Println("Please provide a pokemon name to inspect.")
				continue
			}
			pokemonName := input[1]
			err := commandInspect(pokemonName)
			if err != nil {
				fmt.Println(err)
			}
		case "pokedex":
			if err := commandPokedex(); err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Printf("Unknown command: %s\n", input[0])
		}

	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`
Welcome to the Pokedex!

Usage:

help: Displays a help message
exit: Exits the Pokedex
	`)
	return nil
}

func commandMap() error {
	locationAreas, err := fetchPokemonLocationAreas(false)
	if err != nil {
		return err
	}
	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb() error {
	locationAreas, err := fetchPokemonLocationAreas(true)
	if err != nil {
		return err
	}
	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(areaName string) error {
	fmt.Printf("Exploring %s...\n", areaName)
	encounters, err := fetchLocationAreaData(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandPokedex() error {
	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}

	names := make([]string, 0, len(pokedex))
	for n := range pokedex {
		names = append(names, n)
	}
	sort.Strings(names)

	fmt.Println("Your Pokedex:")
	for _, n := range names {
		fmt.Printf(" - %s\n", n)
	}
	return nil
}
