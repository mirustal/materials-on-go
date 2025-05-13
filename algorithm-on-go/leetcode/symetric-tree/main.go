package symetric_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symetricTree(root.Left, root.Right)
}

func symetricTree(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	return left.Val == right.Val && symetricTree(left.Left, right.Right) && symetricTree(left.Right, right.Left)
}
