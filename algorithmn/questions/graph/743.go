package graph

import "container/heap"

type edge struct {
	from, to, weight int
}

// https://www.youtube.com/watch?v=EaphyqKU4PQ
func NetworkDelayTime(times [][]int, n int, k int) int {
	adjacentList := make(map[int][]edge)
	isVisited := make([]bool, n+1)
	t := 0
	minHeap := make(EdgeMinHeap, 0)
	heap.Init(&minHeap)

	for _, v := range times {
		newEdge := edge{
			from:   v[0],
			to:     v[1],
			weight: v[2],
		}

		v, ok := adjacentList[newEdge.from]
		if ok {
			v = append(v, newEdge)
			adjacentList[newEdge.from] = v
		} else {
			newList := make([]edge, 0)
			newList = append(newList, newEdge)
			adjacentList[newEdge.from] = newList
		}
	}
	heap.Push(&minHeap, edge{
		from:   k,
		to:     k,
		weight: 0,
	})

	for len(minHeap) > 0 {
		current := heap.Pop(&minHeap).(edge)
		if isVisited[current.to] {
			continue
		}

		t = max(t, current.weight)
		isVisited[current.to] = true
		for _, neighbour := range adjacentList[current.to] {
			if !isVisited[neighbour.to] {
				heap.Push(&minHeap, edge{
					from:   current.to,
					to:     neighbour.to,
					weight: neighbour.weight + current.weight,
				})
			}
		}
	}
	for i := 1; i < len(isVisited); i++ {
		if !isVisited[i] {
			return -1
		}
	}
	return t
}

func max(t1 int, t2 int) int {
	if t1 > t2 {
		return t1
	}
	return t2
}

type EdgeMinHeap []edge

func (h EdgeMinHeap) Len() int           { return len(h) }
func (h EdgeMinHeap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h EdgeMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeMinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(edge))
}

func (h *EdgeMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
