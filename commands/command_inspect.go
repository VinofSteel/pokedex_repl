package commands

import (
	"errors"
	"fmt"
)

func Inspect(cfg *Config, param *string) error {
	if param == nil || *param == "" {
		return errors.New("you need to declare the name parameter to inspect a pokémon, like this: inspect <name>")
	}

	pokemon, ok := cfg.CaughtPokemon[*param]
	if !ok {
		return errors.New("you haven't caught this pokémon")
	}

	fmt.Printf("Inspecting %s...\n", pokemon.Name)
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, statStuct := range pokemon.Stats {
		fmt.Printf("-%s: %v\n", statStuct.Stat.Name, statStuct.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeStruct := range pokemon.Types {
		fmt.Printf("-%s\n", typeStruct.Type.Name)
	}

	return nil
}
