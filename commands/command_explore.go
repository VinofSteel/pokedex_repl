package commands

import "fmt"

func Explore(cfg *Config, param *string) error {
	locationResp, err := cfg.PokeapiClient.ExploreLocation(param)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *param)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
