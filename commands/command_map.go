package commands

import (
	"fmt"

	"github.com/vinofsteel/pokedex-repl/client"
)

func Map(config *Config) error {
	fmt.Println("I am the map!")
	client.GetLocations()

	return nil
}
