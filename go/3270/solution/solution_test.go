package solution_test

import (
	"leetcode-go/go/3270/solution"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	output := solution.FindDisappearedNumbers(input)
	expected := []int{5, 6}

	for i := range output {
		if output[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, output)
		}
	}

	input = []int{1, 1}
	output = solution.FindDisappearedNumbers(input)
	expected = []int{2}

	if output[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected, output)
	}

}
