package main

import (
	"bufio"
	"fmt"
	"os"
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
