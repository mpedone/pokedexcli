package main

import (
	"errors"
	"fmt"

	"github.com/mpedone/pokedexcli/internal/pokeapi"
)

func inspect(links *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Yes, the ground is lovely, but please select a pokemon to inspect.")
	}

	if len(args) > 1 {
		return errors.New("Let's take this one thing at a time. Enter the name of ONE pokemon")
	}

	name := args[0]
	if pokemon, ok := pokeapi.Pokedex[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		fmt.Printf("  -hp: %v\n", pokemon.Stats[0].BaseStat)
		fmt.Printf("  -attack: %v\n", pokemon.Stats[1].BaseStat)
		fmt.Printf("  -defense: %v\n", pokemon.Stats[2].BaseStat)
		fmt.Printf("  -special-attack: %v\n", pokemon.Stats[3].BaseStat)
		fmt.Printf("  -special-defense: %v\n", pokemon.Stats[4].BaseStat)
		fmt.Printf("  -speed: %v\n", pokemon.Stats[5].BaseStat)
		fmt.Println("Types:")
		for i := range pokemon.Types {
			fmt.Printf("  - %v\n", pokemon.Types[i].Type.Name)
		}
	}
	return nil
}
