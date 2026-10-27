package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	if entry, ok := c.cache.Get(url); ok {
		pokeResp := Location{}
		err := json.Unmarshal(entry, &pokeResp)
		if err != nil {
			return Location{}, err
		}
		return pokeResp, nil
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

	pokeResp := Location{}
	err = json.Unmarshal(data, &pokeResp)
	if err != nil {
		return pokeResp, err
	}
	c.cache.Add(url, data)
	return pokeResp, nil
}
