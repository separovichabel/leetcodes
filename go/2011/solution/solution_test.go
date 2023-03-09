package solution_test

import (
	"leetcode-go/go/2011/solution"
	"testing"
)

type Case struct {
	input    []string
	expected int
}

var testCases = []Case{
	{
		input:    []string{"--X", "X++", "X++"},
		expected: 1,
	},
	{
		input:    []string{"++X", "++X", "X++"},
		expected: 3,
	},
	{
		input:    []string{"X++", "++X", "--X", "X--"},
		expected: 0,
	},
}

func TestFinalValueAfterOperations(t *testing.T) {
	for _, cas := range testCases {

		output := solution.FinalValueAfterOperations(cas.input)

		if output != cas.expected {
			t.Errorf("For de input \"%v\" it was expected %v but got %v", cas.input, cas.expected, output)
		}
	}
}

func BenchmarkFinalValueAfterOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.FinalValueAfterOperations(cas.input)
		}
	}
}
