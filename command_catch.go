package main

import (
	"errors"
	"fmt"
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
	fmt.Printf("Base exp: %v\n", pDeets.BaseExperience)

	return nil
}
