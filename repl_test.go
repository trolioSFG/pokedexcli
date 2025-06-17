package main

import (
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "	hello  world	",
			expected: []string{"hello", "world"},
		},
		// More cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("actual: %v\n", actual)

		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match got: %d expected: %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Result incorrect, got %s expected %s", word, expectedWord)
			}
		}
	}

	return
}
