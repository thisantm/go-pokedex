package main

import (
	"errors"
	"fmt"
)

func commandExplore(apiState *config, params []string) error {
	if len(params) != 2 {
		return errors.New("a single city name must be passed as a parameter")
	}

	location := params[1]
	resp, err := apiState.pokeapi.GetLocationEncounters(location)

	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Exploring %s\n", location)
	fmt.Println("Found Pokemons:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
