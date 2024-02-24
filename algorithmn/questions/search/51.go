package search

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	result := make([][]string, 0)

	row := ""
	for i := 0; i < n; i++ {
		row = row + "."
	}
	for i := 0; i < n; i++ {
		board[i] = []byte(row)
	}
	backtrack4(n, board, &result, 0, 0)

	return result
}

func backtrack4(n int, board [][]byte, result *[][]string, count int, row int) {
	if count == n {
		temp := make([]string, 0)
		for i := 0; i < n; i++ {
			temp = append(temp, string(board[i]))
		}
		*result = append(*result, temp)
		return
	}

	for j := 0; j < n; j++ {
		if isValid(n, board, row, j) {
			count++
			board[row][j] = "Q"[0]
			backtrack4(n, board, result, count, row+1)
			board[row][j] = "."[0]
			count--
		}
	}

	return
}

func isValid(n int, board [][]byte, x int, y int) bool {
	for i := 0; i < n; i++ {
		if x != i && string(board[i][y]) != "." {
			return false
		}
	}

	for j := 0; j < n; j++ {
		if y != j && string(board[x][j]) != "." {
			return false
		}
	}
	i := x - 1
	j := y - 1
	for i >= 0 && j >= 0 {
		if string(board[i][j]) != "." {
			return false
		}
		i--
		j--
	}

	i = x + 1
	j = y + 1
	for i <= n-1 && j <= n-1 {
		if string(board[i][j]) != "." {
			return false
		}
		i++
		j++
	}

	i = x - 1
	j = y + 1
	for i >= 0 && j <= n-1 {
		if string(board[i][j]) != "." {
			return false
		}
		i--
		j++
	}

	i = x + 1
	j = y - 1
	for i <= n-1 && j >= 0 {
		if string(board[i][j]) != "." {
			return false
		}
		i++
		j--
	}

	return true
}
