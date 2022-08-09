package tree

// RightSideView 二叉树的右视图
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{}
	startIndex := -1
	queue = append(queue, root)
	startIndex++
	res := []int{}

	for len(queue) != startIndex {
		tmpStoreElemLevel := len(queue) - startIndex
		for tmpStoreElemLevel > 0 {
			root = queue[startIndex]
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
			startIndex++
			tmpStoreElemLevel--
		}

		// one level lastNode
		res = append(res, root.Val)
	}
	return res
}
