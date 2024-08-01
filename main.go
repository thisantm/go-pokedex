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

func main() {
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

		command.callback()
	}
}
