package solution_test

import (
	"leetcode-go/go/58/solution"
	"testing"
)

type Case struct {
	input    string
	expected int
}

var testCases = []Case{
	{
		input:    "Hello World",
		expected: 5,
	},
	{
		input:    "   fly me   to   the moon  ",
		expected: 4,
	},
	{
		input:    "luffy is still joyboy",
		expected: 6,
	},
}

func TestIsValid(t *testing.T) {
	for _, cas := range testCases {

		output := solution.LengthOfLastWord(cas.input)

		if output != cas.expected {
			t.Errorf("output %v != %v. input: %v", output, cas.expected, cas.input)
			return
		}
	}
}

func TestIsValid2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.LengthOfLastWord2(cas.input)

		if output != cas.expected {
			t.Errorf("output %v != %v. input: %v", output, cas.expected, cas.input)
			return
		}
	}
}

func BenchmarkLengthOfLastWord2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.LengthOfLastWord2(cas.input)
		}
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.LengthOfLastWord(cas.input)
		}
	}
}
