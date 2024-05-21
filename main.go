package main

import (
	"time"

	"github.com/vinofsteel/pokedex-repl/commands"
	"github.com/vinofsteel/pokedex-repl/internal/client"
)

func main() {
	newClient := client.NewClient(10*time.Second, 10*time.Minute)

	cfg := &commands.Config{
		CaughtPokemon: map[string]client.Pokemon{},
		PokeapiClient: newClient,
	}

	startRepl(cfg)
}
