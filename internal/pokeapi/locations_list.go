package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (respLocations, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	entry, ok := c.Cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return respLocations{}, err
		}

		res, err := c.HttpClient.Do(req)
		if err != nil {
			return respLocations{}, err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return respLocations{}, err
		}

		locations := respLocations{}
		err = json.Unmarshal(data, &locations)
		if err != nil {
			return locations, err
		}
		c.Cache.Add(url, data)
		return locations, nil
	}

	locations := respLocations{}

	err := json.Unmarshal(entry, &locations)
	if err != nil {
		return locations, err
	}

	return locations, nil

}
