package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// HasPathSum
// 二叉树路经总和
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}

// TreePathSum 路径和路径记录
// 二叉树路径总和
func TreePathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	singleRound := make([]int, 0)
	var rangeTree func(root *TreeNode, targetSum int)
	rangeTree = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && targetSum == root.Val {
			singleRound = append(singleRound, root.Val)
			res = append(res, append([]int{}, singleRound...))
			singleRound = singleRound[:len(singleRound)-1]
			return
		}
		singleRound = append(singleRound, root.Val)
		rangeTree(root.Left, targetSum-root.Val)
		rangeTree(root.Right, targetSum-root.Val)
		singleRound = singleRound[:len(singleRound)-1]

	}
	rangeTree(root, targetSum)
	return res
}
