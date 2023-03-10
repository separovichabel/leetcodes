package solution_test

import (
	"leetcode-go/go/1313/solution"
	"testing"
)

type Case struct {
	input    []int
	expected []int
}

var testCases = []Case{
	{
		input:    []int{1, 2, 3, 4},
		expected: []int{2, 4, 4, 4},
	},
	{
		input:    []int{1, 1, 2, 3},
		expected: []int{1, 3, 3},
	},
}

func TestDecompressRLElist(t *testing.T) {
	for _, cas := range testCases {

		output := solution.DecompressRLElist(cas.input)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkDecompressRLElist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.DecompressRLElist(cas.input)
		}
	}
}
