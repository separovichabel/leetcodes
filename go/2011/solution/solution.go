package solution

func FinalValueAfterOperations(operations []string) (x int) {
	for _, str := range operations {
		if str[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return
}
