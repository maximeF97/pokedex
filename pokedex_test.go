package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCleanInputPokedexExamples(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "your butt is well",
			expected: []string{"your", "butt", "is", "well"},
		},
		{
			input:    "I like penuts",
			expected: []string{"i", "like", "penuts"},
		},
		{
			input:    "Big titi wompus",
			expected: []string{"big", "titi", "wompus"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println("actual:", actual)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("input %q: expected %v, got %v", c.input, c.expected, actual)
		}
	}
}
