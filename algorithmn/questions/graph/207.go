package graph

func canFinish(numCourses int, prerequisites [][]int) bool {
	result := make([]int, 0)
	inDegree := make(map[int]int)
	for _, v := range prerequisites {
		_, ok := inDegree[v[0]]
		if ok {
			inDegree[v[0]] = inDegree[v[0]] + 1
		} else {
			inDegree[v[0]] = 1
		}
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		v, ok := inDegree[i]
		if !ok || v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		course := queue[0]
		result = append(result, course)
		queue = queue[1:]
		for _, v := range prerequisites {
			if v[1] == course {
				inDegree[v[0]] = inDegree[v[0]] - 1
				if inDegree[v[0]] == 0 {
					queue = append(queue, v[0])
				}
			}
		}
	}

	if len(result) == numCourses {
		return true
	}
	return false
}
