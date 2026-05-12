package main

import (
	"time"

	"github.com/mpedone/pokedexcli/internal/pokeapi"
)

func main() {
	newClient := pokeapi.PokeClient(5 * time.Second)
	links := &config{
		pokeClient: newClient,
	}
	startRepl(links)
	// newClient := pokeapi.PokeClient(5 * time.Second)
	// pokeapi.PokeReq(&links.client)
}
