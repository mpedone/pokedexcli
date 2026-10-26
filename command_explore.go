package main

import (
	"errors"
	"fmt"
)

// type respPokemon struct {
// 	Region struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"region"`
// }

func commandExplore(links *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	location := args[0]
	pokemon, err := links.pokeClient.ListPokemon(location)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", pokemon.Name)
	fmt.Println("Found Pokemon: ")
	for _, pok := range pokemon.PokemonEncounters {
		fmt.Printf(" - %s\n", pok.Pokemon.Name)
	}

	return nil

	/*url := "https://pokeapi.co/api/v2/location-area/" + location

	entry, err := http.Get(url)

	if err != nil {
		return err*/
	/* return respPokemon{}, err
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respPokemon{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return respPokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return respPokemon{}, err
	}

	pokemons := respPokemon{}
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return pokemons, err
	}
	// c.Cache.Add(url, data)
	return pokemons, nil */
	// }

	/*var pokemons map[string]any
	decoder := json.NewDecoder(entry.Body)
	err = decoder.Decode(&pokemons)
	if err != nil {
		return err
	}

	fmt.Println(pokemons["Region"])*/

	return nil

}
