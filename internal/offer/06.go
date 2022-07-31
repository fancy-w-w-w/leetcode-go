package offer

// 剑指 Offer 06. 从尾到头打印链表
func ReversePrint(head *ListNode) []int {
	stack := []int{}

	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	res := []int{}
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}

	return res
}
