package linkedlist

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
// Partition
func Partition(head *ListNode, x int) *ListNode {
	dummyHeadLess := &ListNode{
		Next: nil,
	}
	dummyHeadMore := &ListNode{
		Next: nil,
	}
	workNodeLess, workNodeMore := dummyHeadLess, dummyHeadMore
	workNode := head
	for workNode != nil {
		nextNode := workNode.Next
		if workNode.Val < x {
			workNodeLess.Next = workNode
			workNodeLess = workNodeLess.Next
		} else {
			workNodeMore.Next = workNode
			workNodeMore = workNodeMore.Next
		}
		workNode = nextNode
	}
	workNodeMore.Next = nil
	workNodeLess.Next = dummyHeadMore.Next
	return dummyHeadLess.Next
}
