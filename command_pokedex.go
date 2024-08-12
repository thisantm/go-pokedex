package main

import (
	"fmt"
)

func commandPokedex(apiState *config, params []string) error {
	if len(apiState.caughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemon, try catching it using the command 'catch {pokemon name}'")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range apiState.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
