package tree

// 把二叉搜索树转换为累加树
// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
// 迭代 反前序做法
func convertBST(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	sum := 0
	head := root
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += root.Val
		root.Val = sum
		root = root.Left
	}
	return head
}
