func groupAnagrams(strs []string) [][]string {
	myMap := make(map[[26]int][]string)
	var masterSlice [][]string 

	for _, value := range strs {
		var key [26]int
		for _, val := range value {
			key[val-'a']++
		}

		myMap[key] = append(myMap[key], value)
	}

	for _, value := range myMap {
		masterSlice = append(masterSlice, value)
	}

	return masterSlice
}
