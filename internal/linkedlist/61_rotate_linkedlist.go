package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// RotateRight 向右翻转链表
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	workNode := head
	listLen := 0
	var rearNode *ListNode
	for workNode != nil {
		listLen++
		// 先记录尾节点
		if workNode.Next == nil {
			rearNode = workNode
		}
		workNode = workNode.Next

	}

	// 取余
	needRotateNum := k % listLen
	// 提前返回
	if needRotateNum == 0 {
		return head
	}

	startMoveIndex := listLen - needRotateNum

	dummyHead := &ListNode{Next: head}
	workNode = head

	var preRearNode *ListNode
	for startMoveIndex > 0 {
		preRearNode = workNode
		workNode = workNode.Next
		startMoveIndex--
	}

	// 从startIndex开始平移后面的整个节点
	rearNode.Next = head
	preRearNode.Next = nil
	dummyHead.Next = workNode

	return dummyHead.Next
}
