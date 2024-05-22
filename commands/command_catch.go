package commands

import (
	"fmt"
	"math/rand"
)

func Catch(cfg *Config, param *string) error {
	pokemon, err := cfg.PokeapiClient.CatchPokemon(param)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a pokéball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.CaughtPokemon[pokemon.Name] = pokemon
	fmt.Println("You may now inspect it with the 'inspect' command.")
	return nil
}
