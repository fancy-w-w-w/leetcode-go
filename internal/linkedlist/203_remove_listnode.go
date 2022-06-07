package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}

	preNode := dummyHead
	for head != nil {
		if head.Val == val {
			head = head.Next
			preNode.Next = head
		} else {
			head = head.Next
			preNode = preNode.Next
		}

	}
	return dummyHead.Next
}
