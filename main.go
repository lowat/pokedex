package main

import (
	"time"

	"github.com/lowat/pokedex/internal/pokeAPI"
)

func main() {
	pokeClient := pokeAPI.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeAPI.Pokemon{},
	}

	startRepl(cfg)
}
