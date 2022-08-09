package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BuildTree 前序和中序构造二叉树
// 构建二叉树
func BuildTreeV1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	var index int = 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}

	root.Left = BuildTreeV1(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = BuildTreeV1(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}

// 中序和后序建立二叉树
// 构建二叉树
// BuildTreeV2
func BuildTreeV2(inorder []int, postorder []int) *TreeNode {
	indexStore := make(map[int]int, len(inorder))

	for i, value := range inorder {
		indexStore[value] = i
	}

	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		// 🈚️剩余节点
		if inorderLeft > inorderRight {
			return nil
		}
		// 取最后节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		inorderIndex := indexStore[val]
		// 先右子节点
		root.Right = build(inorderIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}
