package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mpedone/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocations(pageUrl *string, locCache *pokecache.Cache) (respLocations, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	data := []byte{}

	entry, ok := locCache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return respLocations{}, err
			// log.Fatal(err)
			// return condition, fmt.Errorf("NewRequest err: %w\n", err)
		}

		res, err := c.HttpClient.Do(req)
		if err != nil {
			return respLocations{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return respLocations{}, err
		}
		locCache.Add(url, data)
	} else {
		data = entry
	}

	// Add to the cache
	// map[string]newCache = {time.Now(), data}
	// key = pageURl

	locations := respLocations{}

	err := json.Unmarshal(data, &locations)
	if err != nil {
		return locations, err
	}

	return locations, nil

}
