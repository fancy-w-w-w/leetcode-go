package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// longestConsecutive二叉树最长连续序列
// 给你一棵指定的二叉树的根节点 root ，请你计算其中 最长连续序列路径 的长度
func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 1
	var dfs func(preVal int, root *TreeNode, currentPath int)
	dfs = func(preVal int, root *TreeNode, currentPath int) {
		if preVal+1 != root.Val {
			currentPath = 1
		} else {
			currentPath++
		}
		if root.Left != nil {
			dfs(root.Val, root.Left, currentPath)
		}
		if root.Right != nil {
			dfs(root.Val, root.Right, currentPath)
		}
		if currentPath > res {
			res = currentPath
		}
	}
	dfs(0, root, 0)
	return res
}
