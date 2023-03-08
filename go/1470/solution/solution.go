package solution

// TODO: bit manipulations is also possible
func Shuffle(nums []int, n int) []int {
	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i], nums[i+n])
	}
	return ans
}

func Shuffle2(nums []int, n int) []int {
	ans := make([]int, n*2)

	for i := 0; i < n; i++ {
		ans[i*2], ans[i*2+1] = nums[i], nums[i+n]
	}

	return ans
}
