package main

import (
	"fmt"
	"interview/algorithmn/questions/graph"
)

func main() {
	bytes := []byte{123,
		34,
		100,
		97,
		116,
		97,
		34,
		58,
		123,
		34,
		115,
		104,
		111,
		117,
		108,
		100,
		95,
		99,
		111,
		108,
		108,
		101,
		99,
		116,
		95,
		98,
		105,
		108,
		108,
		105,
		110,
		103,
		95,
		97,
		100,
		100,
		114,
		101,
		115,
		115,
		34,
		58,
		102,
		97,
		108,
		115,
		101,
		44,
		34,
		115,
		104,
		111,
		117,
		108,
		100,
		95,
		99,
		111,
		108,
		108,
		101,
		99,
		116,
		95,
		119,
		57,
		34,
		58,
		102,
		97,
		108,
		115,
		101,
		44,
		34,
		115,
		104,
		111,
		117,
		108,
		100,
		95,
		99,
		111,
		108,
		108,
		101,
		99,
		116,
		95,
		119,
		57,
		95,
		116,
		97,
		120,
		34,
		58,
		102,
		97,
		108,
		115,
		101,
		44,
		34,
		115,
		104,
		111,
		117,
		108,
		100,
		95,
		115,
		104,
		111,
		119,
		95,
		98,
		105,
		108,
		108,
		105,
		110,
		103,
		95,
		97,
		100,
		100,
		114,
		101,
		115,
		115,
		34,
		58,
		102,
		97,
		108,
		115,
		101,
		44,
		34,
		115,
		104,
		111,
		117,
		108,
		100,
		95,
		115,
		104,
		111,
		119,
		95,
		119,
		57,
		95,
		116,
		97,
		120,
		95,
		103,
		117,
		105,
		100,
		101,
		34,
		58,
		102,
		97,
		108,
		115,
		101,
		44,
		34,
		116,
		111,
		116,
		97,
		108,
		95,
		97,
		110,
		110,
		117,
		97,
		108,
		95,
		101,
		120,
		99,
		104,
		97,
		110,
		103,
		101,
		95,
		117,
		115,
		100,
		95,
		99,
		101,
		110,
		116,
		115,
		34,
		58,
		48,
		125,
		44,
		34,
		101,
		120,
		116,
		114,
		97,
		34,
		58,
		123,
		34,
		110,
		111,
		119,
		34,
		58,
		49,
		55,
		48,
		57,
		53,
		54,
		51,
		51,
		52,
		55,
		50,
		53,
		49,
		125,
		44,
		34,
		115,
		116,
		97,
		116,
		117,
		115,
		95,
		99,
		111,
		100,
		101,
		34,
		58,
		48,
		125}

	fmt.Println(string(bytes))
	fmt.Println(graph.MinCostConnectPoints([][]int{{0, 1}, {0, 2}}))
	//fmt.Println(graph.MinCostConnectPoints([][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}}))

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
