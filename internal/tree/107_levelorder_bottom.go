package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 从底层开始
// 层序遍历
func lLevelOrderBottom(root *TreeNode) [][]int {
	//定义一个存储结果的二维切片
	var res [][]int
	//定义一个含有根节点的队列
	arr := []*TreeNode{root}
	//根节点为空直接返回空
	if root == nil {
		return res
	}
	//当队列为空就返回结果，不为空就进入循环
	for len(arr) > 0 {
		//重置队列长度
		size := len(arr)
		//定义一个切片来存储出队列的值
		curRes := []int{}
		//遍历队列
		for i := 0; i < size; i++ {
			//临时变量存储队列中的元素
			node := arr[i]
			//将队列中元素的值加入到这个切片里。
			curRes = append(curRes, node.Val)
			//寻找队列里元素的左右孩子，如果不为空就入队
			if node.Left != nil {
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}
		//遍历队列结束后，把上一层对应的元素出队。
		arr = arr[size:]
		//把所有结果加入到最后的结果上
		res = append(res, curRes)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
