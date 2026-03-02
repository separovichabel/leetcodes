package solution

func KidsWithCandies(candies []int, extraCandies int) []bool {
	greatest := 0
	for _, c := range candies {
		if c > greatest {
			greatest = c
		}
	}

	ans := make([]bool, len(candies))
	for i, c := range candies {
		if (c + extraCandies) >= greatest {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
