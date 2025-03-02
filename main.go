package main

import (
	"strings"
	"time"

	"github.com/arturogood17/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	if len(text) <= 0 {
		return []string{}
	}
	words := strings.Fields(strings.ToLower(text))
	return words
}

func main() {
	c := pokeapi.NewClient(5 * time.Second)
	startREPL(c)
}
