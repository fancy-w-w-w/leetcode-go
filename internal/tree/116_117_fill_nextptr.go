package tree

// 包含Next指针的Node结构
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ConnectLevelNode 层序遍历增加Next指针
func ConnectLevelNode(root *Node) *Node {

	if root == nil {
		return root
	}

	// 初始化队列同时将第一层节点加入队列中，即根节点
	queue := []*Node{root}

	// 循环迭代的是层数
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		// 遍历这一层的所有节点
		for i, node := range tmp {
			// 连接
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			// 拓展下一层节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	// 返回根节点
	return root
}

// O(1) 空间
func ConnectLevelNodeV2(root *Node) *Node {
	if root == nil {
		return root
	}

	// 每次循环从该层的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := leftmost; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right

			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	// 返回根节点
	return root
}

// 带空节点（非完全二叉树）
func ConnectLevelNodeV3(root *Node) *Node {
	if root == nil {
		return nil
	}

	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			// every round nextStart->first node
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}

		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}

		// 到下一层
		start = nextStart

	}
	return root
}
