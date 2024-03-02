package trie

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	Children []*TrieNode
	IsWord   bool // whether there is a word ending here
}

func NewNode() *TrieNode {
	node := TrieNode{
		Children: make([]*TrieNode, 26), // 26 characters
		IsWord:   false,
	}
	return &node
}
func Constructor() Trie {
	return Trie{
		root: NewNode(),
	}
}

func (this *Trie) Insert(word string) {
	currentNode := this.root
	for _, v := range word {
		index := v - 'a'

		if currentNode.Children[index] == nil {
			currentNode.Children[index] = NewNode()
		}
		currentNode = currentNode.Children[index]
	}
	currentNode.IsWord = true
}

func (this *Trie) Search(word string) bool {
	currentNode := this.root
	for _, v := range word {
		index := v - 'a'
		if currentNode.Children[index] == nil {
			return false
		}
		currentNode = currentNode.Children[index]
	}
	return currentNode.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this.root
	for _, v := range prefix {
		index := v - 'a'
		if currentNode.Children[index] == nil {
			return false
		}
		currentNode = currentNode.Children[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
