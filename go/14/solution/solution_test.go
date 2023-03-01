package solution_test

import (
	"leetcode-go/go/14/solution"
	"testing"
)

type Case struct {
	input    []string
	expected string
}

var testCases = []Case{
	{
		input:    []string{"flower", "flow", "flight"},
		expected: "fl",
	},
	{
		input:    []string{"dog", "racecar", "car"},
		expected: "",
	},
}

func TestIsIsomorphic(t *testing.T) {
	for _, cas := range testCases {

		output := solution.LongestCommonPrefix(cas.input)

		if output != cas.expected {
			t.Errorf("output %v != %v", output, cas.expected)
			return
		}
	}

}
