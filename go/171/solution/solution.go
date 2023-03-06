package solution

func TitleToNumber(columnTitle string) int {
	var sum int = 0
	for _, char := range columnTitle {
		sum = sum*26 + int(char-64)
	}
	return sum
}
