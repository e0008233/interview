package main

import (
	"fmt"
	"interview/algorithmn/questions/graph"
)

func main() {

	fmt.Println(graph.FindOrder(2, [][]int{{1, 0}}))
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
