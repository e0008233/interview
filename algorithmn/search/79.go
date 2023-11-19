package search

var direction2 = []int{-1, 0, 1, 0, -1}

func exist(board [][]byte, word string) bool {

	visited := make([][]bool, len(board))
	current := make([]byte, 0)
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			result := backtrack3(board, word, visited, current, i, j)
			if result {
				return true
			}
		}
	}
	return false
}

func backtrack3(board [][]byte, word string, visited [][]bool, current []byte, x, y int) bool {
	if len(current) >= len(word) {
		return string(current) == word
	}
	if visited[x][y] {
		return false
	}

	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0])-1 {
		return false
	}

	visited[x][y] = true
	current = append(current, board[x][y])
	for i := 0; i < 4; i++ {
		found := backtrack3(board, word, visited, current, x+direction2[i], y+direction[i+1])
		if found {
			return true
		}
	}
	current = current[:len(current)-1]
	visited[x][y] = false
	return false
}
