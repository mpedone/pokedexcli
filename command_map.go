package main

import (
	"errors"
	"fmt"
)

func commandMap(links *config) error {
	// if links.next == nil
	// create NewCache
	// else just send the existing cache
	locations, err := links.pokeClient.ListLocations(links.next)
	if err != nil {
		return err
	}

	links.next = locations.Next
	links.prev = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc)
	}

	return nil
}

func commandMapb(links *config) error {
	if links.prev == nil {
		return errors.New("you're on the first page")
	}

	// cache.get(links.prev)
	locations, err := links.pokeClient.ListLocations(links.prev)
	if err != nil {
		return err
	}

	links.next = locations.Next
	links.prev = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc)
	}

	return nil
}
