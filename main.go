package main

import (
	"strings"
)

func cleanInput(text string) []string {
	if len(text) <= 0 {
		return []string{}
	}
	words := strings.Fields(strings.ToLower(text))
	return words
}

func main() {
	startREPL()
}
