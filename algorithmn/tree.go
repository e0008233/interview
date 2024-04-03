package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverse(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	temp := node.Right
	node.Right = node.Left
	node.Left = temp
	reverse(node.Right)
	reverse(node.Left)
	return node
}
