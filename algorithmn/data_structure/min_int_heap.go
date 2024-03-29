package data_structure

import (
	"container/heap"
	"fmt"
	"sort"
)

// An IntMinHeap is a min-heap of ints.
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Test() {

	myHeap := IntMinHeap{3, 2, 1, 4, 5, 6}
	sort.Sort(&myHeap)
	//for _, v := range myHeap {
	//	fmt.Println(v)
	//}
	fmt.Println(myHeap)

	heap.Init(&myHeap)
	fmt.Println(heap.Pop(&myHeap))
	fmt.Println(heap.Pop(&myHeap))

}
