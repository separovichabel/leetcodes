package solution

// O(n)
func NumIdenticalPairs(nums []int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && i < j; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}

	return count
}

// O(n)
func NumIdenticalPairs2(nums []int) int {
	count := 0

	a := make([]int, 101)

	for i := 0; i < len(nums); i++ {
		count = count + a[nums[i]]
		a[nums[i]] = a[nums[i]] + 1
	}

	return count
}

func NumIdenticalPairs3(nums []int) int {
	count := 0

	a := make([]int, 101)

	for i := 0; i < len(nums); i++ {
		a[nums[i]] = a[nums[i]] + 1
	}

	for _, n := range a {
		count = count + ((n * (n - 1)) / 2)
	}

	return count
}
