package google

import "fmt"

func SolveSudoku(board [][]byte) {
	if backtrack2(board) {
		fmt.Println(board)
	}
}

func backtrack2(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := '1'; k <= '9'; k++ {
					isValid := isValidSudoku3(board, i, j, byte(k))
					board[i][j] = byte(k)
					if isValid {
						if backtrack2(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					} else {
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValidSudoku3(board [][]byte, row int, col int, b byte) bool {
	for i := 0; i < 9; i++ {

		if board[i][col] == b {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[row][i] == b {
			return false
		}
	}
	r := row / 3 * 3
	c := col / 3 * 3

	for i := r; i < r+3; i++ {
		for j := c; j < c+3; j++ {
			if board[i][j] == b {
				return false
			}
		}
	}
	return true

}

func isValidSudoku2(board [][]byte) bool {

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
