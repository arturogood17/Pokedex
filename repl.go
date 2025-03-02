package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arturogood17/pokedex/internal/pokeapi"
)

type config struct {
	pokeClient  pokeapi.Client
	nextURL     *string
	previousURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		command := words[0]
		v, ok := getCommands()[command]
		if ok {
			err := v.callback(cfg) //the callback func always returns an err you have to catch
			if err != nil {
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
	}
	return commands
}
