package linkedlist

// 删除链表的倒数第N个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	workNode := head
	node := head
	prenode := dummyHead
	for workNode != nil && n > 0 {
		workNode = workNode.Next
		n--
	}

	for workNode != nil {
		prenode = prenode.Next
		node = node.Next
		workNode = workNode.Next
	}
	prenode.Next = node.Next
	return dummyHead.Next
}
