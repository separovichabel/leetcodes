package solution

func SearchInsert(nums []int, target int) int {
	high := len(nums)
	low := 0
	mid := 0

	for low < high {
		mid = (low + high) / 2

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
	// return sort.SearchInts(nums, target)
}
