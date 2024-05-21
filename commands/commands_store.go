package commands

import "github.com/vinofsteel/pokedex-repl/internal/client"

type Config struct {
	PokeapiClient    client.Client
	nextLocationsURL *string
	prevLocationsURL *string
	CaughtPokemon    map[string]client.Pokemon
}
type CliCommand struct {
	name        string
	description string
	Callback    func(*Config, *string) error // Needs to be exported
}

func GetAll() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			Callback:    MapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			Callback:    MapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area by name and see all the pokemon in it",
			Callback:    Explore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon by name",
			Callback:    Catch,
		},
	}
}
