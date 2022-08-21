package specialquestion

import "project1/internal/linkedlist"

// 大数相减
// 1->2->3  -   2->9
func SubtractionByLinkedlist(list1, list2 *linkedlist.ListNode) *linkedlist.ListNode {
	if list1 == nil || list2 == nil {
		return nil
	}

	workNode1, workNode2 := list1, list2
	stack1, stack2 := []int{}, []int{}
	for workNode1 != nil {
		stack1 = append(stack1, workNode1.Val)
		workNode1 = workNode1.Next
	}
	for workNode2 != nil {
		stack2 = append(stack2, workNode2.Val)
		workNode2 = workNode2.Next
	}
	var res *linkedlist.ListNode

	jiewei := 0
	for len(stack1) != 0 || len(stack2) != 0 || jiewei != 0 {
		var a, b int
		if len(stack1) != 0 {
			a = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			b = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		if a-b-jiewei < 0 {
		}
		cur, isless := absreduce(a - b - jiewei)
		jiewei = isless
		curNode := &linkedlist.ListNode{
			Val:  cur,
			Next: res,
		}
		res = curNode

		// 跳出
		if len(stack1) == 0 && len(stack2) == 0 && jiewei == 1 {
			break
		}
	}
	if jiewei != 0 {
		node := &linkedlist.ListNode{Val: -1, Next: res}
		res = node
	}
	return res

}

func absreduce(a int) (int, int) {
	if a < 0 {
		return 10 + a, 1
	}
	return a, 0
}
