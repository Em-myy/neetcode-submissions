type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var byteSlice []byte
	specialChar := byte('#')

	for _, value := range strs {
		byteSlice = append(byteSlice, byte(len(value)))

		byteSlice = append(byteSlice, specialChar)

		byteSlice = append(byteSlice, value...)
	}

	middleString := string(byteSlice)

	return middleString
}

func (s *Solution) Decode(encoded string) []string {
	var stringSlice []string
	i := 0

	for i < len(encoded) {
		length := int(encoded[i])

		i++
		if encoded[i] != '#' {
			break
		}
		i++

		stringSlice = append(stringSlice, encoded[i:i+length])

		i += length
	}
	return stringSlice
}
