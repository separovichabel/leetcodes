package solution

func DecompressRLElist(nums []int) []int {
	temp := []int{}

	for i := 0; i < len(nums)/2; i++ {
		freq := nums[i*2]
		val := nums[i*2+1]

		for j := 0; j < freq; j++ {
			temp = append(temp, val)
		}
	}

	return temp
}
