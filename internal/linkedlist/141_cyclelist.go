package linkedlist

// 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	dummyHead := &ListNode{
		Next: head,
	}
	preNode, postNode := dummyHead, dummyHead.Next
	for preNode != postNode {
		if postNode == nil || postNode.Next == nil {
			return false
		}
		preNode = preNode.Next
		postNode = postNode.Next.Next
	}
	return true
}
