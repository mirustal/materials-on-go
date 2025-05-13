package balanced_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	return checkDeepth(root) != -1
}

func checkDeepth(node *TreeNode) int {

	if node == nil {
		return 0
	}

	leftDeepth := checkDeepth(node.Left)
	if leftDeepth == -1 {
		return -1
	}

	rightDeepth := checkDeepth(node.Right)
	if rightDeepth == -1 {
		return -1
	}

	if abs(leftDeepth-rightDeepth) > 1 {
		return -1
	}

	return max(leftDeepth, rightDeepth) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
