package search

var direction = []int{-1, 0, 1, 0, -1}

func pacificAtlantic(heights [][]int) [][]int {

	row := len(heights)
	column := len(heights[0])

	pathOneCanReached := make([][]bool, len(heights))
	pathTwoCanReached := make([][]bool, len(heights))

	for i := 0; i < len(heights); i++ {
		pathOneCanReached[i] = make([]bool, len(heights[0]))
		pathTwoCanReached[i] = make([]bool, len(heights[0]))
	}

	for i := 0; i < column; i++ {
		dfs(heights, pathOneCanReached, 0, i)
		dfs(heights, pathTwoCanReached, row-1, i)
	}

	for i := 0; i < row; i++ {
		dfs(heights, pathOneCanReached, i, 0)
		dfs(heights, pathTwoCanReached, i, column-1)
	}

	for i := 0; i < column; i++ {

	}
	result := make([][]int, 0)
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if pathOneCanReached[i][j] && pathTwoCanReached[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func dfs(heights [][]int, reached [][]bool, i, j int) {
	if reached[i][j] {
		return
	}
	reached[i][j] = true

	for x := 0; x < 4; x++ {
		newI := i + direction[x]
		newJ := j + direction[x+1]

		if newI >= 0 && newJ >= 0 && newI < len(heights) && newJ < len(heights[0]) && heights[newI][newJ] >= heights[i][j] {
			dfs(heights, reached, newI, newJ)
		}
	}
}
