package graph

func FindShortestCycle(n int, edges [][]int) int {
	adjacentList := make(map[int]map[int]bool)
	result := -1
	for _, v := range edges {
		value, ok := adjacentList[v[0]]
		if !ok {
			value = make(map[int]bool)
		}
		value[v[1]] = true
		adjacentList[v[0]] = value

		value, ok = adjacentList[v[1]]
		if !ok {
			value = make(map[int]bool)
		}
		value[v[0]] = true
		adjacentList[v[1]] = value
	}

	for _, v := range edges {
		value0, _ := adjacentList[v[0]]
		value0[v[1]] = false
		adjacentList[v[0]] = value0

		value1, _ := adjacentList[v[1]]
		value1[v[0]] = false
		adjacentList[v[1]] = value1

		count, found := bfs(adjacentList, v[0], v[1], n)
		count = count + 1
		if found {
			if result == -1 {
				result = count
			} else {
				if count < result {
					result = count
				}
			}
		}

		value0[v[1]] = true
		adjacentList[v[0]] = value0
		value1[v[0]] = true
		adjacentList[v[1]] = value1
	}

	return result
}

func bfs(list map[int]map[int]bool, i1 int, i2 int, n int) (int, bool) {
	queue := make([]int, 0)
	queue = append(queue, i1)
	steps := -1

	isVisited := make([]bool, n)
	for len(queue) > 0 {
		length := len(queue)
		steps = steps + 1
		for i := 0; i < length; i++ {
			curr := queue[0]
			queue = queue[1:]
			if isVisited[curr] {
				continue
			}

			isVisited[curr] = true
			if curr == i2 {
				return steps, true
			}

			for k, v := range list[curr] {
				if v {
					queue = append(queue, k)
				}
			}
		}
	}
	return steps, false

}
