package main

import (
	"time"

	"github.com/vinofsteel/pokedex-repl/commands"
	"github.com/vinofsteel/pokedex-repl/internal/client"
)

func main() {
	client := client.NewClient(10 * time.Second)
	cfg := &commands.Config{
		PokeapiClient: client,
	}

	startRepl(cfg)
}
