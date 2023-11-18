package main

import (
	"interview/algorithmn/search"
)

func main() {
	println(search.FindCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
}
