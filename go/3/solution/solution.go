package solution

// Given a string s, find the length of the longest substring without duplicate characters.

// ## Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

// ## Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// ## Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// # Constraints:
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.

//  R
//  L
// "abcabcbb"
func LengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	left := 0
	largestLen := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		position, wasSeen := lastSeen[char]

		if wasSeen {
			left = lastSeen[char] + 1
			lastSeen[char] = position
			continue
		}

		lastSeen[char] = right

		if right-left+1 > largestLen {
			largestLen++
		}
	}

	return largestLen
}
