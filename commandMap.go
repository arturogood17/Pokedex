package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func MapCommand(c *config) error {
	var url string
	if c.nextURL == "" {
		url = baseURL + "location-area/"
	} else {
		url = c.nextURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := c.Clnt.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var locationList LocationArea
	json.Unmarshal(data, &locationList)
	c.previousURL = locationList.Previous
	c.nextURL = locationList.Next

	for _, loc := range locationList.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func MapbCommand(c *config) error {
	if c.previousURL == nil {
		return errors.New("you're on the first page")
	}
	url := c.previousURL
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return err
	}

	res, err := c.Clnt.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var locationList LocationArea
	json.Unmarshal(data, &locationList)
	c.previousURL = locationList.Previous
	c.nextURL = locationList.Next

	for _, loc := range locationList.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
