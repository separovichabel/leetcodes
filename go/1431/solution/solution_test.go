package solution_test

import (
	"leetcode-go/go/1431/solution"
	"testing"
)

type Case struct {
	input    []int
	input2   int
	expected []bool
}

var testCases = []Case{
	{
		input:    []int{2, 3, 5, 1, 3},
		input2:   3,
		expected: []bool{true, true, true, false, true},
	},
	{
		input:    []int{4, 2, 1, 1, 2},
		input2:   1,
		expected: []bool{true, false, false, false, false},
	},
	{
		input:    []int{12, 1, 12},
		input2:   10,
		expected: []bool{true, false, true},
	},
}

func TestKidsWithCandies(t *testing.T) {
	for _, cas := range testCases {

		output := solution.KidsWithCandies(cas.input, cas.input2)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkKidsWithCandies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.KidsWithCandies(cas.input, cas.input2)
		}
	}
}
