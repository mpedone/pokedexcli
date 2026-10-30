package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/mpedone/pokedexcli/internal/pokeapi"
)

func catch(links *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("What are you trying to catch, a cold?")
	}

	name := args[0]

	if _, ok := pokeapi.Pokedex[name]; ok {
		return errors.New("Already caught that one!")
	}

	pDeets, err := links.pokeClient.PokeDetails(name)

	if pDeets.BaseExperience < 36 {
		return errors.New("Pokemon doesn't exist... Yet.")
	}

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	exp := pDeets.BaseExperience
	target := int(((25 * math.Log(float64(exp))) - 65) * 0.2)
	roll := rand.IntN(20)

	fmt.Printf(" > Base exp: %v\n", pDeets.BaseExperience)
	fmt.Printf(" > Target: %v\n", target+1)
	fmt.Printf(" > Roll: %v\n", roll+1)

	if roll >= target {
		fmt.Printf("%s caught!\n", name)
		pokeapi.Pokedex[name] = pDeets
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
