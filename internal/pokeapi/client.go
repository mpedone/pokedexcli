package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	HttpClient http.Client
}

func PokeClient(t time.Duration) Client {
	return Client{
		HttpClient: http.Client{
			Timeout: t,
		},
	}
}
