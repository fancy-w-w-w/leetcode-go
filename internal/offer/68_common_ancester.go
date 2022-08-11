package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉树的最近公共祖先
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	leftNode := LowestCommonAncestor(root.Left, p, q)
	rightNode := LowestCommonAncestor(root.Right, p, q)

	if leftNode != nil && rightNode != nil {
		return root
	}

	if leftNode != nil {
		return leftNode
	}

	if rightNode != nil {
		return rightNode
	}

	return nil
}
