package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
)

func catch(links *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("What are you trying to catch, a cold?")
	}

	name := args[0]

	pDeets, err := links.pokeClient.PokeDetails(name)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a ball at %s...\n", name)

	exp := pDeets.BaseExperience
	target := int(((25 * math.Log(float64(exp))) - 65) * 0.2)
	roll := rand.IntN(20)

	fmt.Printf("Base exp: %v\n", pDeets.BaseExperience)
	fmt.Printf("Target: %v\n", target)
	fmt.Printf("Roll: %v\n", roll)

	if roll >= target {
		fmt.Printf("%s caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
