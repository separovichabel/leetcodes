package solution_test

import (
	"leetcode-go/go/205/solution"
	"testing"
)

type Case struct {
	input    []string
	expected bool
}

var testCases = []Case{
	{input: []string{"egg", "add"}, expected: true},
	{input: []string{"foo", "bar"}, expected: false},
	{input: []string{"paper", "title"}, expected: true},
}

func TestIsIsomorphic(t *testing.T) {
	var output bool

	for _, cas := range testCases {

		output = solution.IsIsomorphic(cas.input[0], cas.input[1])

		if output != cas.expected {
			t.Errorf("imput: %v, Expected %v, got %v", cas.input, cas.expected, output)
		}
	}

}
