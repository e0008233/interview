package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}

	ans := make([]*TreeNode, 0)
	root = helper(root, to_delete, &ans)
	if root != nil {
		ans = append(ans, root)
	}
	return ans
}

func helper(root *TreeNode, toDelete []int, ans *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Right = helper(root.Right, toDelete, ans)
	root.Left = helper(root.Left, toDelete, ans)
	for _, v := range toDelete {
		if v == root.Val {
			if root.Right != nil {
				*ans = append(*ans, root.Right)
			}
			if root.Left != nil {
				*ans = append(*ans, root.Left)
			}
			return nil
		}
	}
	return root
}
