package main

import (
	"fmt"
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
	fmt.Println("Hello, World!")
}
