package solution_test

import (
	"leetcode-go/go/1365/solution"
	"testing"
)

type Case struct {
	input    []int
	expected []int
}

var testCases = []Case{
	{
		input:    []int{8, 1, 2, 2, 3},
		expected: []int{4, 0, 1, 1, 3},
	},
	{
		input:    []int{6, 5, 4, 8},
		expected: []int{2, 1, 0, 3},
	},
	{
		input:    []int{7, 7, 7, 7},
		expected: []int{0, 0, 0, 0},
	},
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	for _, cas := range testCases {

		output := solution.SmallerNumbersThanCurrent(cas.input)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkSmallerNumbersThanCurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.SmallerNumbersThanCurrent(cas.input)
		}
	}
}

func TestSmallerNumbersThanCurrent2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.SmallerNumbersThanCurrent2(cas.input)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkSmallerNumbersThanCurrent2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.SmallerNumbersThanCurrent2(cas.input)
		}
	}
}
