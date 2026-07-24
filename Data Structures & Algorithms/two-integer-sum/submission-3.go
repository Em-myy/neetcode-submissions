func twoSum(nums []int, target int) []int {
	/*
	myMap := make(map[int]int)
	for indx, val := range nums {
		myMap[indx] = val;
	}
	for indx1, val1 := range myMap {
		for indx2, val2 := range myMap {
			if indx1 == indx2 {
				continue
			}
			if val1 + val2 == target {
				if indx1 > indx2 {
					return []int{indx2, indx1}
				} else {
					return []int{indx1, indx2}
				}
			}
		}
	}
	return nil
	*/
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				if i > j {
					return []int{j, i}
				} else {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}
