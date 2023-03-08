package solution_test

import (
	"leetcode-go/go/35/solution"
	"testing"
)

type Case struct {
	input1   []int
	input2   int
	expected int
}

var testCases = []Case{
	{
		input1:   []int{1, 3, 5, 6},
		input2:   5,
		expected: 2,
	},
	{
		input1:   []int{1, 3, 5, 6},
		input2:   2,
		expected: 1,
	},
	{
		input1:   []int{1, 3, 5, 6},
		input2:   7,
		expected: 4,
	},
}

func TestSearchInsert(t *testing.T) {
	for _, cas := range testCases {

		output := solution.SearchInsert(cas.input1, cas.input2)

		if output != cas.expected {
			t.Errorf("case: %v %v output %v != %v", cas.input1, cas.input2, output, cas.expected)
			return
		}
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.SearchInsert(cas.input1, cas.input2)
		}
	}
}
