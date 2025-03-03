package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageURL *string) (LocationArea, error) {
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationList LocationArea
	if v, exists := c.cache.Get(url); exists {
		if err := json.Unmarshal(v, &locationList); err != nil {
			return locationList, err
		}
		return locationList, nil
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
	if err := json.Unmarshal(data, &locationList); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)

	return locationList, nil
}
