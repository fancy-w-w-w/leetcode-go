package tree

// 二叉树中所有距离为k的结点
// 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。

// 返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。

func DistanceK(root, target *TreeNode, k int) (ans []int) {
	// 从 root 出发 DFS，记录每个结点的父结点
	parents := map[*TreeNode]*TreeNode{}
	var findParents func(*TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right] = node
			findParents(node.Right)
		}
	}
	findParents(root)

	// 从 target 出发 DFS，寻找所有深度为 k 的结点
	var findAns func(*TreeNode, *TreeNode, int)
	findAns = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k { // 将所有深度为 k 的结点的值计入结果
			ans = append(ans, node.Val)
			return
		}
		// 往下走
		// 判断防止回路
		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		// 防止回路
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}

		// 向上遍历 模拟图
		if parents[node] != from {
			findAns(parents[node], node, depth+1)
		}
	}
	findAns(target, nil, 0)
	return
}
