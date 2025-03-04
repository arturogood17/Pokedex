package main

import (
	"fmt"
	"os"
)

func ExitCommand(c *config, cmds ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
