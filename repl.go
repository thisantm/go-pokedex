package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	readyInput := strings.Fields(output)
	return readyInput
}

func startRepl(apiState *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommandMap := getCommands()
	for {
		print("pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command, ok := cliCommandMap[input[0]]
		if !ok {
			fmt.Println("Command does not exist")
			continue
		}

		err := command.callback(apiState)

		if err != nil {
			fmt.Println(err)
		}
	}
}
