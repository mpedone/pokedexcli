package pokeapi

import (
	"net/http"
	"time"

	"github.com/mpedone/pokedexcli/internal/pokecache"
)

type Client struct {
	Cache      pokecache.Cache
	HttpClient http.Client
}

func PokeClient(t time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(5 * time.Second),
		HttpClient: http.Client{
			Timeout: t,
		},
	}
}
