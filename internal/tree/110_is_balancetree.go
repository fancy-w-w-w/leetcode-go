package tree

import "math"

func IsBalancedTree(root *TreeNode) bool {
	res := true
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := dfs(root.Left) + 1
		rightHeight := dfs(root.Right) + 1

		if math.Abs(float64(leftHeight)-float64(rightHeight)) > float64(1) {
			res = false
		}
		return max(leftHeight, rightHeight)
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
