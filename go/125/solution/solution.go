package solution

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune{}

	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
			runes = append(runes, char)
		}
	}

	i, j := 0, len(runes)-1

	for i < j {
		if runes[i] != runes[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func IsPalindrome2(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		for !IsAlphaNum(s[i]) && i < j {
			i++
		}

		for !IsAlphaNum(s[j]) && i < j {
			j--
		}
		if unicode.ToLower(rune(s[j])) != unicode.ToLower(rune(s[i])) {
			return false
		}

		i++
		j--
	}

	return true
}

func IsAlphaNum(r byte) bool {
	return r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r >= 'A' && r <= 'Z'
}
