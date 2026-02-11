package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
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
