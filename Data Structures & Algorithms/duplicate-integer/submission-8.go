func hasDuplicate(nums []int) bool {
    /* Brute Force Approach
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                return true
            } 
        }
    }
    return false
    */
    numSlice := make(map[int]bool)

    for _, i := range nums {
        if numSlice[i] {
            return true
        }
        numSlice[i] = true
    }
    return false
}
