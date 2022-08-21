package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 两数相加2
// 链表
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 链表的两数之和，考虑栈的使用
	stack1 := []int{}
	stack2 := []int{}
	// 两个链表遍历的val加到栈中
	for l1 != nil || l2 != nil {
		if l1 != nil {
			stack1 = append(stack1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			stack2 = append(stack2, l2.Val)
			l2 = l2.Next
		}
	}
	carry := 0
	var head *ListNode
	var v1, v2 int
	for len(stack1) > 0 || len(stack2) > 0 {
		v1, v2 = 0, 0
		if len(stack1) > 0 {
			v1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			v2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		sum := v1 + v2 + carry
		sum, carry = sum%10, sum/10
		head = &ListNode{Val: sum, Next: head}

	}
	if carry > 0 {
		head = &ListNode{Val: carry, Next: head}
	}
	return head
}
