package solution_test

import (
	"leetcode-go/go/1512/solution"
	"testing"
)

type Case struct {
	input    []int
	expected int
}

var testCases = []Case{
	{
		input:    []int{1, 2, 3, 1, 1, 3},
		expected: 4,
	},
	{
		input:    []int{1, 1, 1, 1},
		expected: 6,
	},
	{
		input:    []int{1, 2, 3},
		expected: 0,
	},
}

func TestNumIdenticalPairs(t *testing.T) {
	for _, cas := range testCases {

		output := solution.NumIdenticalPairs(cas.input)

		if output != cas.expected {
			t.Errorf("For de input []int{%v} it was expected %v but got %v", cas.input, cas.expected, output)
		}

	}
}

func TestNumIdenticalPairs2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.NumIdenticalPairs2(cas.input)

		if output != cas.expected {
			t.Errorf("For de input []int{%v} it was expected %v but got %v", cas.input, cas.expected, output)
		}

	}
}

func TestNumIdenticalPairs3(t *testing.T) {
	for _, cas := range testCases {

		output := solution.NumIdenticalPairs3(cas.input)

		if output != cas.expected {
			t.Errorf("For de input []int{%v} it was expected %v but got %v", cas.input, cas.expected, output)
		}

	}
}

func BenchmarkNumIdenticalPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.NumIdenticalPairs(cas.input)
		}
	}
}

func BenchmarkNumIdenticalPairs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.NumIdenticalPairs2(cas.input)
		}
	}
}

func BenchmarkNumIdenticalPairs3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.NumIdenticalPairs3(cas.input)
		}
	}
}
