package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arturogood17/pokedex/internal/pokeapi"
)

type config struct {
	pokeClient    pokeapi.Client
	nextURL       *string
	previousURL   *string
	caughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		command := words[0]
		var scmd []string
		if len(words) > 1 {
			scmd = words[1:]
		}
		v, ok := getCommands()[command]
		if ok {
			err := v.callback(cfg, scmd...) //the callback func always returns an err you have to catch. You need to add
			if err != nil {                 //the ... for it to accept the commands
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown command: %s\n", command)
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    ExitCommand,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    HelpCommand,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    MapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    MapbCommand,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    ExploreCommand,
		},

		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    CatchCommand,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    InspectCommand,
		},
	}
	return commands
}
