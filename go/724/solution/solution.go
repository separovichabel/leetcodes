package solution

func PivotIndex(nums []int) int {
	var rightSum = 0
	for i := len(nums) - 1; i > -1; i-- {
		rightSum = rightSum + nums[i]
	}

	var leftSum = 0
	var curNum = 0
	for i := 0; i < len(nums); i++ {
		curNum = nums[i]
		rightSum = rightSum - curNum

		if rightSum == leftSum {
			return i
		}

		leftSum = leftSum + curNum
	}
	return -1
}
