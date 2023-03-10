package solution

import "strings"

func MostWordsFound(sentences []string) int {

	max := 0

	for _, c := range sentences {
		words := len(strings.Split(c, " "))
		if words > max {
			max = words
		}
	}

	return max
}

func MostWordsFound2(sentences []string) int {
	max := 0

	for _, sentence := range sentences {
		words := 1

		for _, char := range sentence {
			if char == ' ' {
				words++
			}
		}

		if words > max {
			max = words
		}
	}

	return max
}
