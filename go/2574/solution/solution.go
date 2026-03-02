package solution

func LeftRigthDifference(nums []int) []int {
	leftSum := getLeft(nums)
	rightSum := getRight(nums)

	for i := range rightSum {
		r := leftSum[i] - rightSum[i]
		if r < 0 {
			rightSum[i] = -r
		} else {
			rightSum[i] = r
		}
	}
	return rightSum
}

func getRight(nums []int) []int {
	arr := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			arr[i] = 0
		} else {
			arr[i] = arr[i-1] + nums[i-1]
		}
	}
	return arr
}

func getLeft(nums []int) []int {
	l := len(nums)
	arr := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		if i == l-1 {
			arr[i] = 0
		} else {
			arr[i] = arr[i+1] + nums[i+1]
		}
	}
	return arr
}

func LeftRigthDifference2(nums []int) []int {
	rightSum := 0
	leftSum := 0

	for _, n := range nums {
		rightSum = rightSum + n
	}

	for i, cur := range nums {
		rightSum = rightSum - cur
		nums[i] = abs(leftSum - rightSum)
		leftSum = leftSum + cur
	}

	return nums
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
