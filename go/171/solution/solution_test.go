package solution_test

import (
	"leetcode-go/go/171/solution"
	"testing"
)

type Case struct {
	input    string
	expected int
}

var testCases = []Case{
	{
		input:    "A",
		expected: 1,
	},
	{
		input:    "AB",
		expected: 28,
	},
	{
		input:    "ZY",
		expected: 701,
	},
}

func TestTitleToNumber(t *testing.T) {
	for _, cas := range testCases {

		output := solution.TitleToNumber(cas.input)

		if output != cas.expected {
			t.Errorf("case: %v output %v != %v", cas.input, output, cas.expected)
			return
		}
	}
}

func BenchmarkTitleToNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.TitleToNumber(cas.input)
		}
	}
}
