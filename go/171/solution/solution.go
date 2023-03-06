package solution

import "unicode"

func TitleToNumber(columnTitle string) int {
	var sum int = 0
	for _, r := range columnTitle {
		sum += int(unicode.ToLower(r))
	}
	return sum
}
