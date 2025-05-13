package path_sum_2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	paths := make([]int, 0)

	path(root, &ans, targetSum, paths)

	return ans
}

func path(root *TreeNode, ans *[][]int, sum int, paths []int) {
	if root == nil {
		return
	}

	paths = append(paths, root.Val)

	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
		buf := make([]int, len(paths))
		copy(buf, paths)
		*ans = append(*ans, buf)
		return
	}

	path(root.Left, ans, sum-root.Val, paths)
	path(root.Right, ans, sum-root.Val, paths)
}
