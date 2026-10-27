package main

import (
	"fmt"

	"github.com/mpedone/pokedexcli/internal/pokeapi"
)

func pokedex(links *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range pokeapi.Pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
