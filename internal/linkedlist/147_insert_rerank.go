package linkedlist

var dummyHead_147 *ListNode = &ListNode{}

// InsertionSortList 插入排序进行链表重排
func InsertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	readyList, needSortNode := head, head.Next
	readyList.Next = nil

	for needSortNode != nil {
		nextNode := needSortNode.Next
		readyList = insertToExistList(needSortNode, readyList)
		needSortNode = nextNode
	}
	return readyList
}

func insertToExistList(node *ListNode, head *ListNode) *ListNode {
	preNode, workNode := dummyHead_147, head
	preNode.Next = head
	for workNode != nil {
		if node.Val < workNode.Val {
			node.Next = workNode
			preNode.Next = node
			return dummyHead_147.Next
		}
		preNode = preNode.Next
		workNode = workNode.Next
	}
	// 尾部
	node.Next = nil
	preNode.Next = node
	return dummyHead_147.Next
}
