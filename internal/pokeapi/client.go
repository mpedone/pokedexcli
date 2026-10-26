package pokeapi

import (
	"net/http"
	"time"

	"github.com/mpedone/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func PokeClient(t time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(5 * time.Second),
		httpClient: http.Client{
			Timeout: t,
		},
	}
}
