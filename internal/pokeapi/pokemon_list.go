package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location string) (Location, error) {
	url := baseURL + "/location-area" + "/" + location

	if entry, ok := c.cache.Get(url); ok {
		pokeResp := Location{}
		err := json.Unmarshal(entry, &pokeResp)
		if err != nil {
			return Location{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	pokemons := Location{}
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return pokemons, err
	}
	c.cache.Add(url, data)
	return pokemons, nil

	/* pokemons := Location{}

	err := json.Unmarshal(entry, &pokemons)
	if err != nil {
		return pokemons, err
	}

	return pokemons, nil */
}
