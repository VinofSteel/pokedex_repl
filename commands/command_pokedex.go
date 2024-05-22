package commands

import (
	"errors"
	"fmt"
)

func Pokedex(cfg *Config, param *string) error {
	if len(cfg.CaughtPokemon) == 0 {
		return errors.New("you have no pokémon in your pokédex")
	}

	fmt.Println("Your Pokédex:")
	for k := range cfg.CaughtPokemon {
		fmt.Printf("-%s\n", k)
	}

	return nil
}
