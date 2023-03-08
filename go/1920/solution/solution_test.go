package solution_test

import (
	"leetcode-go/go/1920/solution"
	"testing"
)

type Case struct {
	input    []int
	expected []int
}

var testCases = []Case{
	{
		input:    []int{0, 2, 1, 5, 3, 4},
		expected: []int{0, 1, 2, 4, 5, 3},
	},
	{
		input:    []int{5, 0, 1, 2, 3, 4},
		expected: []int{4, 5, 0, 1, 2, 3},
	},
}

func TestBuildArray(t *testing.T) {
	for _, cas := range testCases {

		output := solution.BuildArray(cas.input)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.BuildArray(cas.input)
		}
	}
}
