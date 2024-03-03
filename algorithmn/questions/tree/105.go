package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	rootIndex := 0
	for i, v := range inorder {
		if v == preorder[0] {
			rootIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:1+rootIndex], inorder[0:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root
}
