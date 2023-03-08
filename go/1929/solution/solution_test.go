package solution_test

import (
	"leetcode-go/go/1929/solution"
	"testing"
)

type Case struct {
	input    []int
	expected []int
}

var testCases = []Case{
	{
		input:    []int{1, 2, 1},
		expected: []int{1, 2, 1, 1, 2, 1},
	},
	{
		input:    []int{1, 3, 2, 1},
		expected: []int{1, 3, 2, 1, 1, 3, 2, 1},
	},
}

func TestGetConcatenation(t *testing.T) {
	for _, cas := range testCases {

		output := solution.GetConcatenation(cas.input)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func TestGetConcatenation2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.GetConcatenation2(cas.input)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkGetConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.GetConcatenation(cas.input)
		}
	}
}

func BenchmarkGetConcatenation2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.GetConcatenation2(cas.input)
		}
	}
}

func BenchmarkGetConcatenation3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.GetConcatenation3(cas.input)
		}
	}
}
