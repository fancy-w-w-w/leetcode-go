package tree

import (
	"math/rand"
	"time"
)

// SortedArrayToBST 将排序数组转换成为二叉平衡树
func SortedArrayToBST(nums []int) *TreeNode {
	rand.Seed(time.Now().Unix())
	var tranverse func(int, int) *TreeNode
	tranverse = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := (left + right + rand.Intn(2)) / 2
		val := nums[mid]
		root := &TreeNode{Val: val}

		root.Left = tranverse(left, mid-1)
		root.Right = tranverse(mid+1, right)

		return root
	}
	return tranverse(0, len(nums)-1)
}
