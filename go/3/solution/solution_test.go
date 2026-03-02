package solution

import "testing"

type testCase struct {
	name     string
	input    string
	expected int
}

var tableTest []testCase = []testCase{
	{
		name:     "Example 1",
		input:    "abcabcbb",
		expected: 3,
	},
	{
		name:     "Example 2",
		input:    "bbbbb",
		expected: 1,
	},
	{
		name:     "Example 3",
		input:    "pwwkew",
		expected: 3,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, c := range tableTest {
		t.Run(c.name, func(t *testing.T) {
			output := LengthOfLongestSubstring(c.input)

			if output != c.expected {
				t.Errorf("Expect %v but got %v", c.expected, output)
			}
		})
	}
}
