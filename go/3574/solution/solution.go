package solution

func SortedSquares(nums []int) []int {
	negatives := []int{}
	positives := []int{}
	result := make([]int, 0, len(nums))

	for _, v := range nums {
		if v < 0 {
			negatives = append(negatives, v*v)
			continue
		}
		positives = append(positives, v*v)
	}

	cursorN := 0
	cursorP := 0

	for i := range result {
		if negatives[cursorN] < positives[cursorP] {
			result[i] = negatives[cursorN]
			cursorN++
			continue
		}
		result[i] = positives[cursorP]
		cursorP++
	}

	return result
}
