package tree

// 二叉搜索树最近公共组先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res := root
	for {
		if res.Val > p.Val && res.Val > q.Val {
			res = res.Left
		} else if res.Val < p.Val && res.Val < q.Val {
			res = res.Right
		} else {
			// 找到了
			return res
		}
	}
	return res
}
