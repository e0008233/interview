package search

func FindCircleNum(grid [][]int) int {
	row := len(grid)

	if row == 0 {
		return 0
	}

	column := len(grid[0])

	isVisited := make([]bool, row)

	queue := make([]int, 0)
	result := 0

	for i := 0; i < row; i++ {
		if isVisited[i] {
			continue
		}
		result++
		queue = append(queue, i)
		for len(queue) > 0 {
			node := queue[0]
			isVisited[node] = true
			queue = queue[1:]
			for j := 0; j < column; j++ {
				if grid[node][j] == 1 && !isVisited[j] {
					queue = append(queue, j)
				}
			}
		}
	}
	return result
}
