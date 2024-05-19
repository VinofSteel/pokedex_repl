package commands

type CliCommand struct {
	name        string
	description string
	Callback    func() error // Needs to be exported
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
	}
}
