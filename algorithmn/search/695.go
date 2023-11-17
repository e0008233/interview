package search

func maxAreaOfIsland(grid [][]int) int {
	row := len(grid)

	if row == 0 {
		return 0
	}

	column := len(grid[0])

	isVisited := make([][]bool, row)
	for i := 0; i < row; i++ {
		isVisited[i] = make([]bool, column)
	}

	queue := make([][]int, 0)
	currentMax := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			result := 0
			if grid[i][j] == 1 && !isVisited[i][j] {
				isVisited[i][j] = true
				queue = append(queue, []int{i, j})
				result++
				for len(queue) > 0 {
					node := queue[0]
					queue = queue[1:]
					if node[0]+1 <= row-1 && !isVisited[node[0]+1][node[1]] && grid[node[0]+1][node[1]] == 1 {
						isVisited[node[0]+1][node[1]] = true
						queue = append(queue, []int{node[0] + 1, node[1]})
						result++
					}

					if node[1]+1 <= column-1 && !isVisited[node[0]][node[1]+1] && grid[node[0]][node[1]+1] == 1 {
						isVisited[node[0]][node[1]+1] = true
						queue = append(queue, []int{node[0], node[1] + 1})
						result++
					}

					if node[0]-1 >= 0 && !isVisited[node[0]-1][node[1]] && grid[node[0]-1][node[1]] == 1 {
						isVisited[node[0]-1][node[1]] = true
						queue = append(queue, []int{node[0] - 1, node[1]})
						result++
					}

					if node[1]-1 >= 0 && !isVisited[node[0]][node[1]-1] && grid[node[0]][node[1]-1] == 1 {
						isVisited[node[0]][node[1]-1] = true
						queue = append(queue, []int{node[0], node[1] - 1})
						result++
					}

				}

				if result > currentMax {
					currentMax = result
				}
			}
		}
	}
	return currentMax
}
