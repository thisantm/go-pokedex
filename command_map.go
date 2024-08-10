package main

import (
	"errors"
	"fmt"
)

func commandMap(apiState *config) error {

	resp, err := apiState.pokeapi.GetLocations(apiState.nextLocationURL)

	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Location Areas:")
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	apiState.nextLocationURL = resp.Next
	apiState.prevLocationURL = resp.Previous

	return nil
}

func commandMapb(apiState *config) error {
	if apiState.prevLocationURL == nil {
		return errors.New("you are on the first page")
	}

	resp, err := apiState.pokeapi.GetLocations(apiState.prevLocationURL)

	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Location Areas:")
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	apiState.nextLocationURL = resp.Next
	apiState.prevLocationURL = resp.Previous

	return nil
}
