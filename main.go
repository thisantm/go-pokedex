package main

import (
	"time"

	"github.com/thisantm/go-pokedex/pokeapi"
)

type config struct {
	pokeapi         pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

func main() {
	apiState := config{
		pokeapi:       pokeapi.NewClient(time.Hour),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(&apiState)
}
