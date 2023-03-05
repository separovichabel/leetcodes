package solution_test

import (
	"leetcode-go/go/20/solution"
	"testing"
)

type Case struct {
	input    string
	expected bool
}

var testCases = []Case{
	{
		input:    "()",
		expected: true,
	},
	{
		input:    "()[]{}",
		expected: true,
	},
	{
		input:    "(]",
		expected: false,
	},
	{
		input:    "",
		expected: true,
	},
	{
		input:    "[()]",
		expected: true,
	},
	{
		input:    "[(])",
		expected: false,
	},
}

func TestIsValid(t *testing.T) {
	for _, cas := range testCases {

		output := solution.IsValid(cas.input)

		if output != cas.expected {
			t.Errorf("output %v != %v. input: %v", output, cas.expected, cas.input)
			return
		}
	}

}
