package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max[T int | float32 | int64 | float64 | uint32 | uint64](t1, t2 T) T {
	if t1 > t2 {
		return t1
	}
	return t2
}
