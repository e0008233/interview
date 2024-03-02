package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter, _ := help(root)
	return diameter
}

func help(node *TreeNode) (int, int) {
	if node == nil {
		return -1, -1
	}
	leftDiameter, leftHeight := help(node.Left)
	rightDiameter, rightHeight := help(node.Right)

	height := max(leftHeight, rightHeight) + 1
	diameter := max(max(leftDiameter, rightDiameter), leftHeight+rightHeight+2)

	return diameter, height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
