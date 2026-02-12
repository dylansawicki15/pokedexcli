package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Show all commands",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Show all areas in a location",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Show all areas in a location, going back one index",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Explore a location area for Pokemon encounters",
		callback:    nil,
	},
	"catch": {
		name:        "catch",
		description: "Catch a Pokemon",
		callback:    nil,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect a caught Pokemon for more details",
		callback:    nil,
	},
	"pokedex": {
 		name:        "pokedex",
 		description: "List your caught Pokemon",
 		callback:    commandPokedex,
 	},
}
