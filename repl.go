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

func startRepl(cfg *commands.Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := sanitizeInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		firstParameter := ""
		if len(words) > 1 {
			firstParameter = words[1]
		}

		command, ok := commands.GetAll()[commandName]
		if !ok {
			fmt.Println("Unknown command, type 'help' to see the available commands!")
			continue
		}

		err := command.Callback(cfg, &firstParameter)
		if err != nil {
			fmt.Println(err)
		}
		continue
	}
}
