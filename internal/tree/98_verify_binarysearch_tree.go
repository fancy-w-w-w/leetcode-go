package tree

import "math"

// IsValidBST 验证是否是二叉搜索树
// 局部最优解，向左寻找时缩短上区间，👉寻找时缩短下区间
// 验证二叉搜索树
func IsValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(root *TreeNode, morethan int, lessthan int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lessthan || root.Val >= morethan {
			return false
		}

		return dfs(root.Left, root.Val, lessthan) && dfs(root.Right, morethan, root.Val)
	}

	return dfs(root, math.MaxInt, math.MinInt)
}

// RecoverBST 恢复BST🌲
// moris算法，维持中序遍历的回溯性
func RecoverBST(root *TreeNode) {
	// x,y 分别为需要交换的节点
	// predcessor是pred的左子树的
	var x, y, pred, predcessor *TreeNode

	for root != nil {
		if root.Left != nil {
			//先找到predcessor节点
			predcessor = root.Left
			for predcessor.Right != nil && predcessor.Right != root {
				predcessor = predcessor.Right
			}

			// 第一次遍历到将predcessor指针指向root
			if predcessor.Right == nil {
				predcessor.Right = root
				root = root.Left
			} else {
				// 左子树已经遍历完成 断开指针链接
				if pred != nil && root.Val < pred.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root
				predcessor.Right = nil
				root = root.Right
			}

		} else {
			// 没有左孩子，不需要找到左子树最大节点
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root
			root = root.Right
		}

	}
	x.Val, y.Val = y.Val, x.Val
}
