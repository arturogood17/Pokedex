package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageURL *string) (LocationArea, error) {
	if v, exists := c.cache.cache[pageURL]; exists {
		return v
	}

	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	var locationList LocationArea
	if err := json.Unmarshal(data, &locationList); err != nil {
		return LocationArea{}, err
	}
	return locationList, nil
}
