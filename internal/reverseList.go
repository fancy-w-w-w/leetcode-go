package internal

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
// 反转指定位置链表
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	needReverseNum := left - right
	if needReverseNum == 0 {
		return head
	}
	dummyHead := &ListNode{Next: head}
	preNode := dummyHead
	for i := left - 1; i > 0; i-- {
		preNode = head.Next
	}
	// init
	workNode := preNode.Next
	for ; needReverseNum > 0; needReverseNum-- {
		tmpNode := workNode.Next
		workNode.Next = tmpNode.Next
		tmpNode.Next = dummyHead.Next
		dummyHead.Next = tmpNode
		fmt.Println(tmpNode.Val)
	}
	return dummyHead.Next
}
