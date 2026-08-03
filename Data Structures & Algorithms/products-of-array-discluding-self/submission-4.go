func productExceptSelf(nums []int) []int {
	preSlice := make([]int, len(nums))
	postSlice := make([]int, len(nums))
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			preSlice[i] = nums[i]
		} else {
			preSlice[i] = nums[i] * preSlice[i - 1]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums) - 1 {
			postSlice[i] = nums[i]
		} else {
			postSlice[i] = nums[i] * postSlice[i + 1]
		}
	}

	for index,_ := range nums {
		if index == 0 {
			result[index] = 1 * postSlice[1]
		} else if index == len(nums) - 1 {
			result[index] = preSlice[index - 1] * 1
		} else {
			result[index] = preSlice[index - 1] * postSlice[index + 1]
		}
	}
	return result
}
