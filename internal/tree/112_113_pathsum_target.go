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
// 二叉树路经和
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

// 路径总和3 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
// dfs解法
func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	// 以每个节点作为根节点
	res := rootSum(root, targetSum)
	// 求根节点子节点是否满足要求
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

// 路径总和3 前缀和+遍历O（n）
// 任意节点、子树
func pathSum3(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)

		// 前缀和满足要求，则加上当前分支上面前缀和出现的次数
		ans += preSum[curr-int64(targetSum)]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
	}
	dfs(root, 0)
	return
}
