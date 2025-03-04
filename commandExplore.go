package main

import (
	"errors"
	"fmt"
)

func ExploreCommand(cfg *config, areas ...string) error {
	if len(areas) == 0 {
		return errors.New("area needed")
	}
	area := areas[0]

	area_poke, err := cfg.pokeClient.AreaPokemon(&area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, val := range area_poke.PokemonEncounters {
		fmt.Printf("- %s\n", val.Pokemon.Name)
	}
	return nil
}
