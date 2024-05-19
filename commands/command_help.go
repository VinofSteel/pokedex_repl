package commands

import "fmt"

func Help() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex REPL CLI!")
	fmt.Println("Usage:")
	fmt.Println()

	// Getting commands from the map where all of the mare stored and showing them to the user
	for _, cmd := range GetAll() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
