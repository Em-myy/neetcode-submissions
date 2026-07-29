func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	result := []int{}

	for _, val := range nums {
		countMap[val]++
	}

	sortSlice := make([]int, 0, len(countMap))

	for _, value := range countMap {
		sortSlice = append(sortSlice, value)
	};

	for i := 0; i < len(sortSlice); i++ {
		for j := 0; j < len(sortSlice) - i - 1; j++ {
			if sortSlice[j] > sortSlice[j+1] {
				sortSlice[j], sortSlice[j+1] = sortSlice[j+1], sortSlice[j]
			}
		}
	}	
	for i := 0; i < k; i++ {
		maxValue := sortSlice[len(sortSlice) - 1 - i]
		for key, value := range countMap {
			if maxValue == value {
				result = append(result, key)
				delete(countMap, key)
				break
			}
		}
	}
	return result
}
