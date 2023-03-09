package solution_test

import (
	"leetcode-go/go/2574/solution"
	"testing"
)

type Case struct {
	input    []int
	expected []int
}

var testCases = []Case{
	{
		input:    []int{10, 4, 8, 3},
		expected: []int{15, 1, 11, 22},
	},
	{
		input:    []int{1},
		expected: []int{0},
	},
}

func TestLeftRigthDifference(t *testing.T) {
	for _, cas := range testCases {

		output := solution.LeftRigthDifference(cas.input)
		firstItem := output[0]

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}

		if firstItem != output[0] {
			t.Errorf("The input slice should not be altered")
		}
	}
}

func TestLeftRigthDifference2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.LeftRigthDifference2(cas.input)

		for i, item := range cas.expected {
			if output[i] != item {
				t.Errorf("For de input []int{%v}[%v] it was expected %v but got %v", cas.input, i, item, output[i])
			}
		}
	}
}

func BenchmarkLeftRigthDifference2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.LeftRigthDifference2(cas.input)
		}
	}
}

func BenchmarkLeftRigthDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.LeftRigthDifference(cas.input)
		}
	}
}
