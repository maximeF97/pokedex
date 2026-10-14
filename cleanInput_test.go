package main

import "testing"

func TestCleanInput(t *testing.T) {
	// Test cases for cleanInput function
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " bIg  BaNane ",
			expected: []string{"big", "banane"},
		},
		{
			input:    " SOo  Stankyy ",
			expected: []string{"soo", "stankyy"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len dont match")
			continue
		}
		for i := range actual {
			world := actual[i]
			expectedWorld := c.expected[i]
			if expectedWorld != world {
				t.Errorf("expected %v got %v", expectedWorld, world)
			}
		}

	}
}
