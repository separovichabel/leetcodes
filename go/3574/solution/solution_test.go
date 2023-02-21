package solution_test

import (
	"leetcode-go/go/3574/solution"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	input := []int{-4, -1, 0, 3, 10}
	output := solution.SortedSquares(input)
	expected := []int{0, 1, 9, 16, 100}

	for i := range input {
		if output[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, output)
		}
	}

	input = []int{-7, -3, 2, 3, 11}
	output = solution.SortedSquares(input)
	expected = []int{4, 9, 9, 49, 121}

	for i := range input {
		if output[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, output)
		}
	}

}
