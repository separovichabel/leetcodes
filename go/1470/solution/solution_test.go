package solution_test

import (
	"leetcode-go/go/1470/solution"
	"testing"
)

type Case struct {
	input    []int
	input2   int
	expected []int
}

var testCases = []Case{
	{
		input:    []int{2, 5, 1, 3, 4, 7},
		input2:   3,
		expected: []int{2, 3, 5, 4, 1, 7},
	},
	{
		input:    []int{1, 2, 3, 4, 4, 3, 2, 1},
		input2:   4,
		expected: []int{1, 4, 2, 3, 3, 2, 4, 1},
	},
	{
		input:    []int{1, 1, 2, 2},
		input2:   2,
		expected: []int{1, 2, 1, 2},
	},
}

func TestShuffle(t *testing.T) {
	for _, cas := range testCases {

		output := solution.Shuffle(cas.input, cas.input2)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func TestShuffle2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.Shuffle2(cas.input, cas.input2)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkShuffle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.Shuffle(cas.input, cas.input2)
		}
	}
}

func BenchmarkShuffle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.Shuffle2(cas.input, cas.input2)
		}
	}
}
