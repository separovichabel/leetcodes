package solution

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TwoSum2(nums []int, target int) []int {
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]

		expctNum := target - cur

		j, ok := m[expctNum]

		m[cur] = i
		if !ok {
			continue
		}

		if j == i {
			continue
		}

		return []int{i, j}
	}
	return []int{}
}
