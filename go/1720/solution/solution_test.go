package solution_test

import (
	"leetcode-go/go/1720/solution"
	"testing"
)

type Case struct {
	input    []int
	input2   int
	expected []int
}

var testCases = []Case{
	{
		input:    []int{1, 2, 3},
		input2:   1,
		expected: []int{1, 0, 2, 1},
	},
	{
		input:    []int{6, 2, 7, 3},
		input2:   4,
		expected: []int{4, 2, 0, 7, 4},
	},
}

func TestDecode(t *testing.T) {
	for _, cas := range testCases {

		output := solution.Decode(cas.input, cas.input2)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.Decode(cas.input, cas.input2)
		}
	}
}
