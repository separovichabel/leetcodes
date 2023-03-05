package solution

import "strings"

func LengthOfLastWord(s string) int {
	f := strings.Fields(s)
	return len(f[len(f)-1])
}

func LengthOfLastWord2(s string) int {
	startsAt := len(s)
	finishesAt := startsAt

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && finishesAt == startsAt {
			finishesAt--
			startsAt--
		} else if s[i] != ' ' {
			startsAt--
		} else {
			break
		}
	}
	return finishesAt - startsAt
}
