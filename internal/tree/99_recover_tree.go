package tree

var last, first, second *TreeNode

// 恢复二叉搜索树
// 找到中序遍历不满足单调递增的位置
func RecoverTree(root *TreeNode) {
	last, first, second = nil, nil, nil
	dfs(root)
	first.Val, second.Val = second.Val, first.Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if last != nil && root.Val <= last.Val {
		if first != nil {
			second = root
			return //剪枝
		}
		first, second = last, root
	}
	last = root
	dfs(root.Right)
}
