package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	registry := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name: 		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
		"explain": {
			name: 		 "explain",
			description: "Explains what this is",
			callback:    commandExplain,
		},
		"map": {
			name: 		 "map forward",
			description: "List next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name: 		 "map back",
			description: "List previous 20 locations",
			callback:    commandMapb,
		},
	}

	links := config{
		next: "https://pokeapi.co/api/v2/location-area/",
		prev: "",
	}

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
		// fmt.Printf("Your command was %s\n", registry[word].name)

		command, ok := registry[word]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&links)
            if err != nil {
                fmt.Println(err)
            }
			if word == "help" {
				for _, desc := range registry {
				    fmt.Printf("%s: %s\n", desc.name, desc.description)
				}
				fmt.Println("")
			}
		}
		// registry[word].callback()
		// switch word {
		// case "exit":
		// 	registry[word].callback()
		
		// }
	}
}