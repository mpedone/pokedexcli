package main

import (
	"time"

	"github.com/mpedone/pokedexcli/internal/pokeapi"
	"github.com/mpedone/pokedexcli/internal/pokecache"
)

func main() {
	newClient := pokeapi.PokeClient(5 * time.Second)
	theCache := pokecache.NewCache(5 * time.Second)
	links := &config{
		pokeClient: newClient,
		pokeCache:  theCache,
	}
	startRepl(links)
}
