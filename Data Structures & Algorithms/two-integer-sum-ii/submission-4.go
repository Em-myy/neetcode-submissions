func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	mySlice := []int{}

	for left < right {
		currentSum := numbers[left] + numbers[right]

		if currentSum > target {
			right--
		} else if currentSum < target {
			left++
		} else {
			mySlice = append(mySlice, left+1, right+1)
			return mySlice
		}
	} 
	return nil
}
