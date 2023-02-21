package solution_test

import (
	"leetcode-go/go/724/solution"
	"testing"
)

func TestPivotIndex1(t *testing.T) {
	input := []int{1, 7, 3, 6, 5, 6}
	output := solution.PivotIndex(input)
	expected := 3

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}

}

func TestPivotIndex2(t *testing.T) {
	input := []int{1, 2, 3}
	output := solution.PivotIndex(input)
	expected := -1

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestPivotIndex3(t *testing.T) {
	input := []int{2, 1, -1}
	output := solution.PivotIndex(input)
	expected := 0

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}
