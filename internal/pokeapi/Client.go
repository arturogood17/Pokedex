package pokeapi

import (
	"net/http"
	"time"

	"github.com/arturogood17/pokedex/internal/pokecache"
)

type Client struct { // http.Client ya tiene timeout
	cache      pokecache.Cache
	HttpClient http.Client
}

func NewClient(timeout, interval time.Duration) Client {
	c := Client{
		cache: pokecache.NewCache(interval),
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
	return c
}
