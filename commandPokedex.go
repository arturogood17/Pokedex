package main

import (
	"errors"
	"fmt"
)

func PokedexCommand(c *config, cmds ...string) error {
	if len(c.caughtPokemon) == 0 {
		return errors.New("no pokemons caught")
	}
	for k := range c.caughtPokemon {
		fmt.Println(" - ", k)
	}
	return nil
}
