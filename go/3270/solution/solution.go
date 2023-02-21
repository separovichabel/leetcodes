package solution

func FindDisappearedNumbers(nums []int) []int {
	var maxNum int = len(nums)
	var res []int = make([]int, maxNum)
	for _, v := range nums {
		if v <= maxNum {
			res[v-1] = 1
		}
	}
	respDefinitiva := []int{}
	for i, v := range res {
		if v == 0 {
			respDefinitiva = append(respDefinitiva, i+1)
		}
	}
	return respDefinitiva
}
