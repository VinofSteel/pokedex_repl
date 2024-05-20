package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vinofsteel/pokedex-repl/commands"
)

func sanitizeInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := sanitizeInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, ok := commands.GetAll()[commandName]
		if !ok {
			fmt.Println("Unknown command, type 'help' to see the available commands!")
			continue
		}

		config := &commands.Config{
			Next:     "placeholder",
			Previous: "placeholder",
		}

		err := command.Callback(config)
		if err != nil {
			fmt.Println(err)
		}
		continue
	}
}
