package tree

// IsSymmetric 是否是对称树
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var dfs func(leftNode *TreeNode, rihtNode *TreeNode) bool
	dfs = func(leftNode *TreeNode, rihtNode *TreeNode) bool {
		if leftNode == nil && rihtNode == nil {
			return true
		} else if leftNode != nil && rihtNode != nil && leftNode.Val == rihtNode.Val {
			return dfs(leftNode.Left, rihtNode.Right) && dfs(leftNode.Right, rihtNode.Left)
		}
		return false
	}
	return dfs(root.Left, root.Right)
}
