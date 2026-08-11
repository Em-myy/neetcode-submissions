func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	};

	numsMap := make(map[int]struct{})

	for _, value := range nums {
		numsMap[value] = struct{}{}
	};

	maxStreak := 0

	for num := range numsMap {
		if _, hasPrevious := numsMap[num - 1]; !hasPrevious {
			currentStreak := 1
			for {
				if _, hasNext := numsMap[num + currentStreak]; hasNext {
					currentStreak++
				} else {
					break
				}
			}
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		}
	}
	return maxStreak
}
