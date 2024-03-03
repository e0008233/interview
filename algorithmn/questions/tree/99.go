package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var prev, isFirstWrong, isSecondWrong *TreeNode

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	prev, isFirstWrong, isSecondWrong = nil, nil, nil
	helper2(root)
	temp := isFirstWrong.Val
	isFirstWrong.Val = isSecondWrong.Val
	isSecondWrong.Val = temp

	return
}

func helper2(root *TreeNode) {
	if root == nil {
		return
	}

	helper2(root.Left)
	if (prev != nil) && prev.Val >= root.Val {
		if isFirstWrong == nil {
			isFirstWrong = prev
		}
		isSecondWrong = root
	}
	prev = root
	helper2(root.Right)

	return
}
