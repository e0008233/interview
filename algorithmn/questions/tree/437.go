package tree

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// Count paths starting from current node
	count := countPaths(root, targetSum)

	// Recursively traverse left and right subtrees
	count += pathSum(root.Left, targetSum)
	count += pathSum(root.Right, targetSum)

	return count
}

func countPaths(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}

	// Check if current node's value matches targetSum
	paths := 0
	if node.Val == targetSum {
		paths++
	}

	// Recursively check paths from left and right subtrees
	paths += countPaths(node.Left, targetSum-node.Val)
	paths += countPaths(node.Right, targetSum-node.Val)

	return paths
}
