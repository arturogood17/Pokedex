package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/arturogood17/pokedex/internal/pokeapi"
)

func CatchCommand(c *config, cmds ...string) error {
	if len(cmds) != 0 {
		return errors.New("provide a pokemon name")
	}
	name := cmds[0]
	pokemon, err := c.pokeClient.PokemonCatch(name)

	if err != nil {
		return err
	}
	var caught = make(map[string]pokeapi.Pokemon)
	bingo := 5 + 1
	catchingProb := rand.Intn(bingo)
	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
	if catchingProb == bingo {
		fmt.Printf("%s was caught\n", pokemon.Name)
		caught[pokemon.Name] = pokemon
	}
	return nil
}
