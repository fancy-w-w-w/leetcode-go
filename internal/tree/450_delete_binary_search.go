package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// DeleteNode删除二叉搜索树节点
// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	var left, right *TreeNode

	if key < root.Val {
		left = DeleteNode(root.Left, key)
		root.Left = left
	} else if key > root.Val {
		right = DeleteNode(root.Right, key)
		root.Right = right
	} else {
		left = root.Left
		right = root.Right
		for right != nil && right.Left != nil {
			right = right.Left
		}
		if right != nil {
			right.Left = left
			return root.Right
		} else {
			return left
		}
	}
	return root
}
