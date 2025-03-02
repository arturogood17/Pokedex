package main

import (
	"errors"
	"fmt"
)

func MapCommand(c *config) error {
	locationList, err := c.pokeClient.ListLocation(c.nextURL)
	if err != nil {
		return err
	}

	for _, loc := range locationList.Results {
		fmt.Println(loc.Name)
	}
	c.nextURL = locationList.Next
	c.previousURL = locationList.Previous

	return nil
}

func MapbCommand(c *config) error {
	if c.previousURL == nil {
		return errors.New("you're on the first page")
	}
	locationList, err := c.pokeClient.ListLocation(c.previousURL)
	if err != nil {
		return err
	}
	c.previousURL = locationList.Previous
	c.nextURL = locationList.Next

	for _, loc := range locationList.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
