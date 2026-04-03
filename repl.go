package main

import (
	"strings"
    "fmt"
    "bufio"
    "os"
    "io"
    "log"
    "net/http"
    "encoding/json"
    // "reflect"
)

func startRepl() {
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
        links := config{
		next: "https://pokeapi.co/api/v2/location-area/",
		prev: "",
	}

		command, ok := getCommands()[word]
		if !ok {
			fmt.Println("Unknown command")
            continue
		} else {
            err := command.callback(&links)
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
    return map[string]cliCommand {
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
			name: 		 "map",
			description: "List next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name: 		 "mapb",
			description: "List previous 20 locations",
			callback:    commandMapb,
		},
    }
}

type config struct {
    next string
    prev string
}

type Keys struct {
    Count    int    `json:"count"`
    Next     string `json:"next"`
    Previous string    `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
}

func commandMap(links *config) error {
    
    fmt.Printf("next: %s, prev: %s\n", links.next, links.prev)

    res, err := http.Get(links.next)
    // res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
    }
    if err != nil {
        log.Fatal(err)
    }

    var entries Keys

    err2 := json.Unmarshal(body, &entries)
     if err2 != nil {
         fmt.Println(err2)
     }

    areas := entries.Results

    links.next = entries.Next
    links.prev = entries.Previous

    for _, x := range areas {
        fmt.Println(x.Name)
    }

    return nil
}

func commandMapb(links *config) error {
    if links.prev == "" {
        fmt.Println("you're on the first page")
        return nil
    }

    links.next = links.prev
    commandMap(links)
    return nil

}
