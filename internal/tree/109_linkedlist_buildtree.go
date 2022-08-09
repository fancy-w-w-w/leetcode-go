package tree

import "project1/internal/linkedlist"

// sortedListToBST 将有序链表转换为二叉平衡搜索树
// 有序链表转换二叉搜索树
func SortedListToBST(head *linkedlist.ListNode) *TreeNode {
	globalHead := head
	var length int
	for ; head != nil; head = head.Next {
		length++
	}
	var buildTree func(int, int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		// 填充空节点
		if left > right {
			return nil
		}
		mid := (left + right + 1) / 2

		// 先不填充root节点，先构建左子树，再填充中间节点
		root := &TreeNode{}
		root.Left = buildTree(left, mid-1)
		root.Val = globalHead.Val
		globalHead = globalHead.Next
		root.Right = buildTree(mid+1, right)
		return root
	}
	return buildTree(0, length-1)
}
