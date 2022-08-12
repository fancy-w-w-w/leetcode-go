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
// 任意时刻，\textit{fast}fast 指针走过的距离都为 \textit{slow}slow 指针的 22 倍。因此，我们有

// a+(n+1)b+nc=2(a+b) \implies a=c+(n-1)(b+c)
// a+(n+1)b+nc=2(a+b)⟹a=c+(n−1)(b+c)

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
