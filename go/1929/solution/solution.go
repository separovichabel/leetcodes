package solution

func GetConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	copy(ans, nums)
	copy(ans[len(nums):], nums)
	return ans
}

func GetConcatenation2(nums []int) []int {
	lenN := len(nums)
	ans := make([]int, lenN*2)
	for i, d := range nums {
		ans[i], ans[i+lenN] = d, d
	}
	return ans
}

func GetConcatenation3(nums []int) []int {
	return append(nums, nums...)
}
