package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonCatch(name string) (Pokemon, error) {
	url := baseURL + "pokemon" + name

	var Poke Pokemon
	if val, exists := c.cache.Get(url); exists {
		if err := json.Unmarshal(val, &Poke); err != nil {
			return Pokemon{}, nil
		}
		return Poke, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if err = json.Unmarshal(data, &Poke); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return Poke, nil
}
