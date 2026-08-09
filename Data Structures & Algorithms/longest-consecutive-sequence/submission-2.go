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
			currentNum := num
			currentStreak := 1

			for {
				if _, hasNext := numsMap[currentNum + 1]; hasNext {
					currentNum++
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
