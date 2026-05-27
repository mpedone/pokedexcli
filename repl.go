package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mpedone/pokedexcli/internal/pokeapi"
	"github.com/mpedone/pokedexcli/internal/pokecache"
	// "reflect"
)

func startRepl(links *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		clean := cleanInput(text)
		if len(clean) == 0 {
			continue
		}
		word := clean[0]
		/* links := &config{
			next: "https://pokeapi.co/api/v2/location-area/",
			prev: "",
		} */

		command, ok := getCommands()[word]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(links)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explain": {
			name:        "explain",
			description: "Explains what this is",
			callback:    commandExplain,
		},
		"map": {
			name:        "map",
			description: "List next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 locations",
			callback:    commandMapb,
		},
	}
}

type config struct {
	pokeClient pokeapi.Client
	pokeCache  pokecache.Cache
	next       *string
	prev       *string
}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
}
