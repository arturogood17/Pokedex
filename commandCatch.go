package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func CatchCommand(c *config, cmds ...string) error {
	if len(cmds) == 0 {
		return errors.New("provide a pokemon name")
	}
	name := cmds[0]
	pokemon, err := c.pokeClient.PokemonCatch(name)

	if err != nil {
		return err
	}
	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
	if res > 40 { //why 40? We'll never know!
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	c.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
