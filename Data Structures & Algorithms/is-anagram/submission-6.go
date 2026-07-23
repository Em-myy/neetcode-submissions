func isAnagram(s string, t string) bool {
	firstMap := make(map[rune]int);
	secondMap := make(map[rune]int);

	for _, val := range s {
		firstMap[val]++
	}
	for _, val := range t {
		secondMap[val]++
	}

	if len(firstMap) != len(secondMap) {
		return false
	}

	for key, val1 := range firstMap {
		val2, exists := secondMap[key]
		if !exists || val1 != val2 {
			return false
		}
	}
	return true
}
