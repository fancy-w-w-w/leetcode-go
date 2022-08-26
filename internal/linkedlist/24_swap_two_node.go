package linkedlist

// 两两交换链表中的节点
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nowNext := head.Next
	head.Next = swapPairs(nowNext.Next)
	nowNext.Next = head
	return nowNext
}
