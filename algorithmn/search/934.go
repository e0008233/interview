package search

var direction3 = []int{-1, 0, 1, 0, -1}

func ShortestBridge(grid [][]int) int {

	isVisited := make([][]bool, len(grid))
	isVisited2 := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		isVisited[i] = make([]bool, len(grid[0]))
		isVisited2[i] = make([]bool, len(grid[0]))

	}
	queue := make([][]int, 0)
	isFound1 := false
	for i := 0; i < len(grid); i++ {
		if isFound1 {
			break
		}
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				bfs(grid, i, j, &queue, isVisited)
				isFound1 = true
			}
			if isFound1 {
				break
			}
		}
	}

	isFound := false
	count := 0

	for _, v := range queue {
		isVisited2[v[0]][v[1]] = true
	}
	for !isFound && len(queue) > 0 {
		num := len(queue)
		for i := 0; i < num; i++ {
			node := queue[0]
			queue = queue[1:]
			for k := 0; k < 4; k++ {
				newI := node[0] + direction3[k]
				newJ := node[1] + direction3[k+1]
				if newI > len(grid)-1 || newI < 0 || newJ < 0 || newJ > len(grid[0])-1 {
					continue
				}
				if isVisited2[newI][newJ] {
					continue
				}
				isVisited2[newI][newJ] = true

				if grid[newI][newJ] == 1 {
					return count
				}

				queue = append(queue, []int{newI, newJ})
			}
		}
		count++

	}
	return count
}

func bfs(grid [][]int, i, j int, result *[][]int, isVisited [][]bool) {

	if i > len(grid)-1 || i < 0 || j < 0 || j > len(grid[0])-1 {
		return
	}
	if isVisited[i][j] {
		return
	}

	if grid[i][j] != 1 {
		return
	}

	isVisited[i][j] = true
	*result = append(*result, []int{i, j})

	for k := 0; k < 4; k++ {
		bfs(grid, i+direction3[k], j+direction3[k+1], result, isVisited)
	}
	return
}
