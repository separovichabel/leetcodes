package solution

func BuildArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i := range nums {
		ans[i] = nums[nums[i]]
	}

	return ans
}
