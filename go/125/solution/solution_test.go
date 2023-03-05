package solution_test

import (
	"leetcode-go/go/125/solution"
	"testing"
)

type Case struct {
	input    string
	expected bool
}

var testCases = []Case{
	{
		input:    "A man, a plan, a canal: Panama",
		expected: true,
	},
	{
		input:    "race a car",
		expected: false,
	},
	{
		input:    " ",
		expected: true,
	},
	{
		input:    "0P",
		expected: false,
	},
	{
		input:    " 0P",
		expected: false,
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, cas := range testCases {

		output := solution.IsPalindrome(cas.input)

		if output != cas.expected {
			t.Errorf("case: %v output %v != %v", cas.input, output, cas.expected)
			return
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.IsPalindrome(cas.input)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.IsPalindrome2(cas.input)

		if output != cas.expected {
			t.Errorf("case: %v output %v != %v", cas.input, output, cas.expected)
			return
		}
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.IsPalindrome2(cas.input)
		}
	}
}
