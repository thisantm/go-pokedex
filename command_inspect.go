package main

import (
	"errors"
	"fmt"
)

func commandInspect(apiState *config, params []string) error {
	if len(params) != 2 {
		return errors.New("a single pokemon name must be passed as a parameter")
	}

	pokemonName := params[1]
	pokemon, ok := apiState.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("you haven't caught that pokemon yet")
	}

	fmt.Printf(`Name: %s
Height: %d
Weight: %d
`,
		pokemon.Name,
		pokemon.Height,
		pokemon.Weight,
	)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}

	return nil
}
