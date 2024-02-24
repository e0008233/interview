package main

import (
	"container/heap"
	"fmt"
	"interview/algorithmn/data_structure"
)

func main() {
	h := &data_structure.IntMaxHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	//fmt.Println(dp.DiffWaysToCompute("2-1-1"))
}
