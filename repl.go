package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mpedone/pokedexcli/internal/pokeapi"
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
		args := []string{}
		if len(clean) > 1 {
			args = clean[1:]
		}

		command, ok := getCommands()[word]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(links, args...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"explore": {
			name:        "explore <location_name>",
			description: "List pokemon in a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Get those pokemon!",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Look at those pokemon!",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List those pokemon!",
			callback:    pokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

type config struct {
	pokeClient pokeapi.Client
	next       *string
	prev       *string
	location   *string
}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
}
