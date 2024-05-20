package commands

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}
type CliCommand struct {
	name        string
	description string
	Callback    func(*Config) error // Needs to be exported
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
			description: "Get the names of 20 Pokémon maps by page",
			Callback:    Map,
		},
	}
}
