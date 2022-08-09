package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 环形链表
// 检测链表是否具有环，并返回环的位置
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var res *ListNode
	var isContainedCycle bool
	preNode, postNode := head, head
	for preNode != nil {
		postNode = postNode.Next
		if preNode.Next == nil {
			return nil
		}
		preNode = preNode.Next.Next
		if preNode == postNode {
			isContainedCycle = true
			res = head
			break
		}
	}
	if isContainedCycle {
		for preNode != res {
			preNode = preNode.Next
			res = res.Next
		}
	}
	return res
}
