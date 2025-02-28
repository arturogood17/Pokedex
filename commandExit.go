package main

import (
	"fmt"
	"os"
)

func ExitCommand() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
