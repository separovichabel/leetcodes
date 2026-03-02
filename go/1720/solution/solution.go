package solution

func Decode(encoded []int, first int) []int {
	arr := make([]int, len(encoded)+1)

	arr[0] = first

	for i, el := range encoded {
		arr[i+1] = arr[i] ^ el
	}

	return arr
}
