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
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		clean := cleanInput(text)
		word := clean[0]
		// fmt.Printf("Your command was %s\n", registry[word].name)

		command, ok := registry[word]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.callback()
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