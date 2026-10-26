package main

import (
	"errors"
	"fmt"
)

// type respPokemon struct {
// 	Region struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"region"`
// }

func commandExplore(links *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	location := args[0]
	pokemon, err := links.pokeClient.ListPokemon(location)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", pokemon.Name)
	fmt.Println("Found Pokemon: ")
	for _, pok := range pokemon.PokemonEncounters {
		fmt.Printf(" - %s\n", pok.Pokemon.Name)
	}

	return nil
}
