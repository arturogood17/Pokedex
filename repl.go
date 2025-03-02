package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arturogood17/pokedex/internal/pokeapi"
)

func startREPL(c *pokeapi.Client) {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{
		Clnt:        c,
		nextURL:     "",
		previousURL: nil,
	}

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

type config struct {
	Clnt        *pokeapi.Client
	nextURL     string
	previousURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
			description: "Shows list of locations",
			callback:    MapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows previous list of locations",
			callback:    MapbCommand,
		},
	}
	return commands
}
