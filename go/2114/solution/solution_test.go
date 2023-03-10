package solution_test

import (
	"leetcode-go/go/2114/solution"
	"testing"
)

type Case struct {
	input    []string
	expected int
}

var testCases = []Case{
	{
		input:    []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
		expected: 6,
	},
	{
		input:    []string{"please wait", "continue to fight", "continue to win"},
		expected: 3,
	},
}

func TestMostWordsFound(t *testing.T) {
	for _, cas := range testCases {

		output := solution.MostWordsFound(cas.input)

		if output != cas.expected {
			t.Errorf("imput: %v, Expected %v, got %v", cas.input, cas.expected, output)
		}
	}

}

func BenchmarkMostWordsFound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.MostWordsFound(cas.input)
		}
	}
}

func TestMostWordsFound2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.MostWordsFound2(cas.input)

		if output != cas.expected {
			t.Errorf("imput: %v, Expected %v, got %v", cas.input, cas.expected, output)
		}
	}

}

func BenchmarkMostWordsFound2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.MostWordsFound2(cas.input)
		}
	}
}
