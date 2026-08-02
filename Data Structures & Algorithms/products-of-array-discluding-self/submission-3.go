func productExceptSelf(nums []int) []int {
	result := []int{}
	for index,_ := range nums {
		innerSlice := []int{}
		product := 1
		for ind, val := range nums {
			if ind != index {
				innerSlice = append(innerSlice, val)
			}
		}
		for _, v := range innerSlice {
			product *= v
		}
		result = append(result, product)
	}
	return result
}
