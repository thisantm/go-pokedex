package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

`)
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("bye bye")
	os.Exit(0)
	return nil
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	readyInput := strings.Fields(output)
	return readyInput
}

func startRepl() {
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
