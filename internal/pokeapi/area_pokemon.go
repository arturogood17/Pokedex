package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) AreaPokemon(area *string) (Area, error) {
	if area == nil {
		return Area{}, errors.New("no valid area provided")
	}

	url := baseURL + "location-area/" + *area

	var PokemonList Area
	if v, exists := c.cache.Get(url); exists {
		if err := json.Unmarshal(v, &PokemonList); err != nil {
			return PokemonList, err
		}
		return PokemonList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Area{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return Area{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Area{}, err
	}
	if err := json.Unmarshal(data, &PokemonList); err != nil {
		return Area{}, err
	}

	c.cache.Add(url, data)

	return PokemonList, nil
}
