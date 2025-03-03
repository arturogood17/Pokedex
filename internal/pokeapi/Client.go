package pokeapi

import (
	"net/http"
	"time"

	"github.com/arturogood17/pokedex/internal/pokecache"
)

type Client struct { // http.Client ya tiene timeout
	HttpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {
	c := Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
	return c
}
