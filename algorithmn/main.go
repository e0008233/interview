package main

import "interview/algorithmn/google"

func main() {
	board := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	byteBoard := make([][]byte, len(board))
	for i, row := range board {
		byteBoard[i] = make([]byte, len(row))
		for j, cell := range row {
			byteBoard[i][j] = cell[0]
		}
	}

	google.SolveSudoku(byteBoard)
	//fmt.Println(google.NextPermutation([]int{1,2,3}))
}
