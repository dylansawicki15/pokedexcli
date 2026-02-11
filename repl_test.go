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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "aasdf asdf  asdf  asdf",
			expected: []string{"aasdf", "asdf", "asdf", "asdf"},
		},
		{
			input:    "1123 4567 89",
			expected: []string{"1123", "4567", "89"},
		},
	}

	caseNum := 0
	for _, c := range cases {
		caseNum++
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("case %d: expected %d words, got %d", caseNum, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("case %d word %d: expected %v, got %v", caseNum, i, expectedWord, word)
			}
		}
	}

}
