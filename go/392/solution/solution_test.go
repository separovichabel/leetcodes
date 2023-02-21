package solution_test

import (
	"leetcode-go/go/392/solution"
	"testing"
)

type Case struct {
	input    []string
	expected bool
}

var testCases = []Case{
	{input: []string{"abc", "ahbgdc"}, expected: true},
	{input: []string{"axc", "ahbgdc"}, expected: false},
	{input: []string{"ah", "ahbgdc"}, expected: true},
}

func TestIsIsomorphic(t *testing.T) {
	var output bool

	for _, cas := range testCases {

		output = solution.IsSubsequence(cas.input[0], cas.input[1])

		if output != cas.expected {
			t.Errorf("imput: %v, Expected %v, got %v", cas.input, cas.expected, output)
		}
	}

}
