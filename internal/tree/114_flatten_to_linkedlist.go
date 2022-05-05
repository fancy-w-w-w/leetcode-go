package tree

// Flatten 将二叉树原地展开为链表
func Flatten(root *TreeNode) {
	cur := root

	for cur != nil {
		if cur.Left != nil {
			// 求左子树最右节点
			next := cur.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			// 左子树最右节点右指针指向根右节点
			pre.Right = cur.Right
			// 旋转
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}
