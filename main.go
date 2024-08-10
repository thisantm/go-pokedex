package main

import "github.com/thisantm/go-pokedex/pokeapi"

type config struct {
	pokeapi         pokeapi.Client
	// nextLocationURL *string
	// prevLocationURL *string
}

func main() {
	apiState := config{
		pokeapi: pokeapi.NewClient(),
	}
	startRepl(&apiState)
}
