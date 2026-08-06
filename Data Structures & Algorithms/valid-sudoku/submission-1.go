func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[int]map[byte]bool)
	colMap := make(map[int]map[byte]bool)
	boxMap := make(map[int]map[byte]bool)

	for rowIndex, value := range board {
		for colIndex, val := range value {
			if val == '.' {
				continue
			}

			box := (rowIndex / 3) * 3 + (colIndex / 3)

			if rowMap[rowIndex] == nil {
				rowMap[rowIndex] = make(map[byte]bool)
			}
			if colMap[colIndex] == nil {
				colMap[colIndex] = make(map[byte]bool)
			}
			if boxMap[box] == nil {
				boxMap[box] = make(map[byte]bool)
			}

			if rowMap[rowIndex][val] || colMap[colIndex][val] || boxMap[box][val] {
				return false
			}
			
			rowMap[rowIndex][val] = true
			colMap[colIndex][val] = true
			boxMap[box][val] = true
		}
	}
	return true
}
