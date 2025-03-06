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

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		pokeClient:    pokeClient,
	}
	startREPL(cfg)
}
