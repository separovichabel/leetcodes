package solution

import "sort"

func SmallerNumbersThanCurrent(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	indexes := map[int]int{}

	for i, n := range sorted {
		if _, ok := indexes[n]; !ok {
			indexes[n] = i
		}
	}

	for i, n := range nums {
		sorted[i] = indexes[n]
	}

	return sorted
}

func SmallerNumbersThanCurrent2(nums []int) []int {
	ocur := map[int]int{}

	for _, n := range nums {
		if _, ok := ocur[n]; !ok {
			ocur[n] = 1
		} else {
			ocur[n]++
		}
	}

	carry := map[int]int{}
	totalCarry := 0

	for i := 0; i <= 100; i++ {
		if ocurCount, ok := ocur[i]; ok {
			carry[i] = totalCarry
			totalCarry = totalCarry + ocurCount
		}
	}

	ans := make([]int, len(nums))
	for i, n := range nums {
		ans[i] = carry[n]
	}

	return ans
}
