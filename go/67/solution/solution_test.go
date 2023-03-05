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
		input:    [2]string{"0", "0"},
		expected: "0",
	},
	{
		input:    [2]string{"00", "001"},
		expected: "001",
	},
	{
		input: [2]string{
			"111",
			"1010",
		},
		expected: "10001",
	},
}

func TestAddBinary(t *testing.T) {
	for _, cas := range testCases {

		output := solution.AddBinary(cas.input[0], cas.input[1])

		if output != cas.expected {
			t.Errorf("case: %v output %v != %v", cas.input, output, cas.expected)
			return
		}
	}
}

func BenchmarkAddBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.AddBinary(cas.input[0], cas.input[1])
		}
	}
}

func TestAddBinary2(t *testing.T) {
	for _, cas := range testCases {

		output := solution.AddBinary2(cas.input[0], cas.input[1])

		if output != cas.expected {
			t.Errorf("case: %v output %v != %v", cas.input, output, cas.expected)
			return
		}
	}
}

func BenchmarkAddBinary2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, cas := range testCases {
			solution.AddBinary2(cas.input[0], cas.input[1])
		}
	}
}
