package main

import (
	"fmt"
	"interview/algorithmn/questions/graph"
)

func main() {

	fmt.Println(graph.FindRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
