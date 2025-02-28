package main

import (
	"bufio"
	"fmt"
	"os"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		command := words[0]
		v, ok := getCommands()[command]
		if ok {
			err := v.callback() //the callback func always returns an err you have to catch
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
	return commands
}
