package tree

import (
	"fmt"
)

// FindDuplicateSubtrees 给定一棵二叉树 root，返回所有重复的子树。
// 如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	// 从叶子节点开始，相同的节点加入
	// 遍历相同节点父节点，如果父节点和其他节点相同，进入下一步
	// 兄弟节点有标记也相同，证明是相同子树
	// 序列化

	res := []*TreeNode{}
	map1 := make(map[string]int, 0)
	var postTraverse func(traNode *TreeNode) string
	postTraverse = func(traNode *TreeNode) string {
		if traNode == nil {
			return ""
		}

		hashVal := fmt.Sprintf("%d.%s.%s", traNode.Val, postTraverse(traNode.Left), postTraverse(traNode.Right))

		if v, ok := map1[hashVal]; ok && v == 1 {
			res = append(res, traNode)
		}
		map1[hashVal]++
		return hashVal
	}

	postTraverse(root)
	return res
}
