package main

import (
	"fmt"
)

func commandHelp(apiState *config, params []string) error {
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
