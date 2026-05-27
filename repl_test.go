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
			input:    "Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  helLO world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello WORLD  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected: %v, got: %v", len(c.expected), len(actual))
			// t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word)
				// t.Fail()
			}
		}
	}
}
