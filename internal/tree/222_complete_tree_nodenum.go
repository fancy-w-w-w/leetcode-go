package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// CountNodes 统计完全二叉树节点个数
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	getLeftHeight := func(root *TreeNode) int {
		hight := 0
		for root != nil {
			hight++
			root = root.Left
		}
		return hight
	}

	getRightHeight := func(root *TreeNode) int {
		hight := 0
		for root != nil {
			hight++
			root = root.Right
		}
		return hight
	}
	left, right := getLeftHeight(root.Left), getRightHeight(root.Right)
	if left == right {
		return (1 << (left + 1)) - 1
	}
	return CountNodes(root.Left) + CountNodes(root.Right) + 1

}
