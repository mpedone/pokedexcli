package main

import "fmt"

func commandExplain(links *config, args ...string) error {
	fmt.Println("")
	fmt.Println("A Pokedex is an index of Pokemon (nevermind what a Pokemon is). This is a CLI-tool to create/access a Pokedex.")
	fmt.Println("")
	return nil
}
