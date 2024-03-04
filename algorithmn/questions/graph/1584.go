package graph

import (
	"container/heap"
)

type pointer struct {
	x, y int
}

type distance struct {
	to       pointer
	distance int
}

func MinCostConnectPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	adjacentList := make(map[pointer][]distance)
	isVisited := make(map[pointer]bool)
	for i := 0; i < len(points); i++ {
		point1 := pointer{
			x: points[i][0],
			y: points[i][1],
		}
		for j := i + 1; j < len(points); j++ {
			point2 := pointer{
				x: points[j][0],
				y: points[j][1],
			}
			distanceNum := abs(point1.x-point2.x) + abs(point1.y-point2.y)
			val, ok := adjacentList[point1]
			if !ok {
				val = make([]distance, 0)
			}
			val = append(val, distance{
				to:       point2,
				distance: distanceNum,
			})
			adjacentList[point1] = val

			val, ok = adjacentList[point2]
			if !ok {
				val = make([]distance, 0)
			}
			val = append(val, distance{
				to:       point1,
				distance: distanceNum,
			})
			adjacentList[point2] = val
		}
	}
	minHeap := make(distanceHeap, 0)
	heap.Init(&minHeap)
	heap.Push(&minHeap, distance{
		to: pointer{
			x: points[0][0],
			y: points[0][1],
		},
		distance: 0,
	})
	result := 0
	count := 0
	for len(minHeap) > 0 {
		curr := heap.Pop(&minHeap).(distance)
		v, ok := isVisited[curr.to]

		if ok && v {
			continue
		}
		isVisited[curr.to] = true
		count++
		result = result + curr.distance
		if count == len(points) {
			return result
		}
		for _, distanceEdge := range adjacentList[curr.to] {
			heap.Push(&minHeap, distanceEdge)
		}

	}

	return result
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return 0 - i
}

type distanceHeap []distance

func (h distanceHeap) Len() int           { return len(h) }
func (h distanceHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h distanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *distanceHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(distance))
}

func (h *distanceHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
