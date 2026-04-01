package main

import (
	"strings"
    "fmt"
    "os"
    "io"
    "log"
    "net/http"
    "encoding/json"
    // "reflect"
)

// func startRepl() {
//     scanner := bufio.NewScanner(os.Stdin)
//     for {
// 		fmt.Print("Pokedex > ")
// 		scanner.Scan()
// 		text := scanner.Text()
// 		clean := cleanInput(text)
//         if len(clean) == 0 {
//             continue
//         }
// 		word := clean[0]
// 		// fmt.Printf("Your command was %s\n", registry[word].name)

// 		command, ok := registry[word]
// 		if !ok {
// 			fmt.Println("Unknown command")
//             continue
// 		} else {
//             err := command.callback(&links)
//             if err != nil {
//                 fmt.Println(err)
//             }
// 			if word == "help" {
// 				for _, desc := range registry {
// 				    fmt.Printf("%s: %s\n", desc.name, desc.description)
// 				}
// 				fmt.Println("")
// 			}
// 		}
// 	}
// }

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

func commandExit(links *config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(links *config) error {
    fmt.Println("")
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println("")
    return nil
} 

func commandExplain(links *config) error {
    fmt.Println("")
    fmt.Println("A Pokedex is an index of Pokemon (nevermind what a Pokemon is). This is a CLI-tool to create/access a Pokedex.")
    fmt.Println("")
    return nil
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
