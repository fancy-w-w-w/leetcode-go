package linkedlist

import "fmt"

// ReverseBetween 翻转指定位置 指定区间的链表
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	needReverseNum := right - left
	if needReverseNum == 0 {
		return head
	}
	dummyHead := &ListNode{Next: head}
	preNode := dummyHead
	for i := left - 1; i > 0; i-- {
		preNode = preNode.Next
	}
	// init
	workNode := preNode.Next
	fmt.Println(workNode.Val)
	for ; needReverseNum > 0; needReverseNum-- {
		tmpNode := workNode.Next
		workNode.Next = tmpNode.Next
		tmpNode.Next = preNode.Next
		preNode.Next = tmpNode
		fmt.Println(tmpNode.Val)
	}
	return dummyHead.Next
}
