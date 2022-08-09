package dynamicresolve

import "project1/internal/tree"

// RobV1 打家劫舍
func RobV1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// 0不偷 1偷
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+nums[i])
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

// 打家劫舍2
func robv2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(RobV1(nums[:len(nums)-1]), RobV1(nums[1:]))
}

// 打家劫舍3
// 二叉树形式
func robV3(root *tree.TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *tree.TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}
