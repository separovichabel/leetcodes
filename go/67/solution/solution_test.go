package solution_test

import (
	"leetcode-go/go/67/solution"
	"testing"
)

type Case struct {
	input    [2]string
	expected string
}

var testCases = []Case{
	{
		input:    [2]string{"11", "1"},
		expected: "100",
	},
	{
		input:    [2]string{"1010", "1011"},
		expected: "10101",
	},
	{
		input:    [2]string{"111", "111"},
		expected: "1110",
	},
	{
		input:    [2]string{"00", "001"},
		expected: "001",
	},
}

func TestIsValid(t *testing.T) {
	for _, cas := range testCases {

		output := solution.AddBinary(cas.input[0], cas.input[1])

		if output != cas.expected {
			t.Errorf("output %v != %v. input: %v", output, cas.expected, cas.input)
			return
		}
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.AddBinary(cas.input[0], cas.input[1])
		}
	}
}
