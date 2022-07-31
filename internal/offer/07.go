package offer

// 剑指 Offer 07. 重建二叉树
// 先序和后序构建二叉树
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootval := preorder[0]
	root := &TreeNode{
		Val: rootval,
	}

	index := 0
	for j := range inorder {
		if inorder[j] == rootval {
			index = j
			break
		}
	}

	root.Left = BuildTree(preorder[1:index+1], inorder[:index])
	root.Right = BuildTree(preorder[index+1:], inorder[index+1:])
	return root
}
