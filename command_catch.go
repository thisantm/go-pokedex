package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(apiState *config, params []string) error {
	if len(params) != 2 {
		return errors.New("a single pokemon name must be passed as a parameter")
	}

	pokemon := params[1]
	resp, err := apiState.pokeapi.GetPokemon(pokemon)

	if err != nil {
		return err
	}

	catchForm := min(95, 95/(0.95+1/(2*math.Pow(10, 4))*math.Pow(float64(resp.BaseExperience), 2)))
	catchRate := int(max(5, catchForm))
	chance := rand.Intn(100)

	fmt.Printf("Trying to catch a %s\n", pokemon)
	if chance >= catchRate {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon)
	apiState.caughtPokemon[pokemon] = resp

	return nil
}
