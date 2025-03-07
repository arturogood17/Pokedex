package main

import (
	"errors"
	"fmt"
)

func InspectCommand(c *config, cmds ...string) error {
	if len(cmds) != 1 {
		return errors.New("provide a valid pokemon name")
	}
	name := cmds[0]
	val, exists := c.caughtPokemon[name]
	if !exists {
		return errors.New("you haven't caught this pokemon")

	}
	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %d\n", val.Height)
	fmt.Printf("Weight: %d\n", val.Weight)
	fmt.Println("Stats:")
	for _, v := range val.Stats {
		fmt.Printf("   -%s: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range val.Types {
		fmt.Printf("   - %s\n", v.Type.Name)
	}
	return nil
}
