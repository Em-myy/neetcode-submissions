func isPalindrome(s string) bool {
	beforeArray := []string{}

	for i := 0; i < len(s); i++ {
		b := s[i]

		if (b >= 'A' && b <= 'Z')  {
			b = b + 32
		}

		if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
			beforeArray = append(beforeArray, string(b))
		}
	}

	left := 0
	right := len(beforeArray) - 1
	for left < right {
		if beforeArray[left] != beforeArray[right] {
			return false
		}
		left++
		right--
	}
	return true
}
