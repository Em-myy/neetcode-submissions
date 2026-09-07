func threeSum(nums []int) [][]int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	outerSlice := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1	
		target := -nums[i]

		for left < right {
			sums := nums[left] + nums[right]

			if sums < target {
				left++
			} else if sums > target {
				right--
			} else {
					outerSlice = append(outerSlice, []int{nums[i], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right --
					}

					left++
					right--
			}
		}
	}
	return outerSlice
}
