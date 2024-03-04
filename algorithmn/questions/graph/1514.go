package graph

import "container/heap"

type edge2 struct {
	to   int
	prob float64
}

func MaxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	adjacentList := make(map[int][]edge2)
	isVisited := make([]bool, n)
	for i, v := range edges {

		val, ok := adjacentList[v[0]]
		edge1 := edge2{
			to:   v[1],
			prob: succProb[i],
		}
		if !ok {
			val = make([]edge2, 0)
			val = append(val, edge1)
			adjacentList[v[0]] = val
		} else {
			val = append(val, edge1)
			adjacentList[v[0]] = val
		}

		val, ok = adjacentList[v[1]]
		edge3 := edge2{
			to:   v[0],
			prob: succProb[i],
		}
		if !ok {
			val = make([]edge2, 0)
			val = append(val, edge3)
			adjacentList[v[1]] = val
		} else {
			val = append(val, edge3)
			adjacentList[v[1]] = val
		}
	}
	maxHeap := make(EdgeMaxHeap, 0)
	heap.Init(&maxHeap)
	heap.Push(&maxHeap, edge2{
		to:   start_node,
		prob: 1,
	})

	for len(maxHeap) > 0 {
		curr := heap.Pop(&maxHeap).(edge2)
		if isVisited[curr.to] {
			continue
		}
		isVisited[curr.to] = true
		if curr.to == end_node {
			return curr.prob
		}

		for _, v := range adjacentList[curr.to] {
			heap.Push(&maxHeap, edge2{
				to:   v.to,
				prob: curr.prob * v.prob,
			})
		}
	}

	return 0

}

type EdgeMaxHeap []edge2

func (h EdgeMaxHeap) Len() int           { return len(h) }
func (h EdgeMaxHeap) Less(i, j int) bool { return h[i].prob > h[j].prob }
func (h EdgeMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeMaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(edge2))
}

func (h *EdgeMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
