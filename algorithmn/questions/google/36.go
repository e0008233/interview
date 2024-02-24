package google

func isValidSudoku(board [][]byte) bool {

	rowMap := make([]map[byte]int, 9)
	for i, _ := range rowMap {
		rowMap[i] = make(map[byte]int)
	}

	columnMap := make([]map[byte]int, 9)
	for i, _ := range columnMap {
		columnMap[i] = make(map[byte]int)
	}
	boxMap := make([]map[byte]int, 9)
	for i, _ := range boxMap {
		boxMap[i] = make(map[byte]int)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			} else {
				_, ok := rowMap[i][board[i][j]]
				if ok {
					return false
				} else {
					rowMap[i][board[i][j]] = 1
				}

				_, ok2 := columnMap[j][board[i][j]]
				if ok2 {
					return false
				} else {
					columnMap[j][board[i][j]] = 1
				}

				_, ok3 := boxMap[i/3+j/3*3][board[i][j]]
				if ok3 {
					return false
				} else {
					boxMap[i/3+j/3*3][board[i][j]] = 1
				}

			}
		}
	}

	return true
}
