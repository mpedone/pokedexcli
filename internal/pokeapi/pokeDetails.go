package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

var Pokedex map[string]PokeDeets = make(map[string]PokeDeets)

func (c *Client) PokeDetails(name string) (PokeDeets, error) {
	url := baseURL + "/pokemon/" + name

	if entry, ok := c.cache.Get(url); ok {
		deets := PokeDeets{}
		err := json.Unmarshal(entry, &deets)
		if err != nil {
			return PokeDeets{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeDeets{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeDeets{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeDeets{}, err
	}

	poke := PokeDeets{}
	err = json.Unmarshal(data, &poke)
	if err != nil {
		return poke, err
	}
	c.cache.Add(url, data)
	return poke, nil
}
