package solution_test

import (
	"leetcode-go/go/1480/solution"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output := solution.RunningSum(input)
	expected := []int{1, 3, 6, 10}

	for i := range output {
		if output[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, output)
		}
	}

	input = []int{1, 1, 1, 1, 1}
	output = solution.RunningSum(input)
	expected = []int{1, 2, 3, 4, 5}

	if output[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected, output)
	}

	input = []int{3, 1, 2, 10, 1}
	output = solution.RunningSum(input)
	expected = []int{3, 4, 6, 16, 17}

	if output[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected, output)
	}

}
