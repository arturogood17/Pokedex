package main

import (
	"errors"
	"testing"
	"time"

	"github.com/arturogood17/pokedex/internal/pokeapi"
)

func TestCommandCatch(t *testing.T) {
	cases := []struct {
		input    string
		expected error
	}{
		{
			input:    "Boltund",
			expected: nil,
		},
		{
			input:    "Nosepass",
			expected: nil,
		},
		{
			input:    "Stantler",
			expected: nil,
		},
		{
			input:    "Cevinho",
			expected: errors.New("provide a pokemon name"),
		},
	}
	c := &config{
		pokeClient: pokeapi.NewClient(5*time.Second, 50*time.Second),
	}
	for index, val := range cases {
		err := CatchCommand(c, val.input)
		if err != nil {
			t.Errorf("Error when running test %d - actual: %v", index, err)
			return
		}
	}
}
