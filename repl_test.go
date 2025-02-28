package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hakuna matata ",
			expected: []string{"hakuna", "matata"},
		},
		{
			input:    "HAKUNA MATATA ",
			expected: []string{"hakuna", "matata"},
		},
		{
			input:    "HAKUNA   mataTA",
			expected: []string{"hakuna", "matata"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected len: %v - Actual len: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Expected word: %v - Actual word: %v", c.expected[i], actual[i])
			}
		}
	}
}
