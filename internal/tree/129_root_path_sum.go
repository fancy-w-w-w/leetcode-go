package tree

func dfs1(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs1(root.Left, sum) + dfs1(root.Right, sum)
}

// 根结点到叶节点数字之和
func sumNumbers(root *TreeNode) int {
	return dfs1(root, 0)
}
