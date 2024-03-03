package main

import (
	"fmt"
	"interview/algorithmn/questions/trie"
)

func main() {
	obj := trie.Constructor()
	obj.Insert("apple")
	fmt.Println(obj)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
