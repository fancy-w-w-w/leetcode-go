package dfs

import (
	"math"
	"project1/internal/tree"
)

// MaxPathSum 二叉树最大路径和
// 从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列
func MaxPathSum(root *tree.TreeNode) int {
	res := math.MinInt
	var dfs func(root *tree.TreeNode) int
	dfs = func(root *tree.TreeNode) int {
		if root == nil {
			return 0
		}
		leftMaxValue := max(dfs(root.Left), 0)
		RightMaxValue := max(dfs(root.Right), 0)

		singleMaxPath := root.Val + leftMaxValue + RightMaxValue
		// 先比较left->root->right路径顺序
		res = max(res, singleMaxPath)

		// 再返回left->root || right->root顺序
		return root.Val + max(leftMaxValue, RightMaxValue)
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
