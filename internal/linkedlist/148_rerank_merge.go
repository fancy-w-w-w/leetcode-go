package linkedlist


// SortLinkedList 链表排序 O(nlogn) 原地
func SortLinkedList(head *ListNode) *ListNode{
	workNode := head
	length := 0
	dummyHead := &ListNode{Next: head}
	for workNode != nil{
		workNode = workNode.Next
		length++
	}
    for subLength := 1; subLength < length; subLength <<= 1 {
        prev, cur := dummyHead, dummyHead.Next
        for cur != nil {
            head1 := cur
            for i := 1; i < subLength && cur.Next != nil; i++ {
                cur = cur.Next
            }

            head2 := cur.Next
            cur.Next = nil
            cur = head2
            for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
                cur = cur.Next
            }

            var next *ListNode
            if cur != nil {
                next = cur.Next
                cur.Next = nil
            }

            prev.Next = mergeReRankList(head1, head2)

            for prev.Next != nil {
                prev = prev.Next
            }
            cur = next
        }
	}
	return dummyHead.Next
}

// mergeReRankList 归并排序局部序列
func mergeReRankList(head1 *ListNode, head2 *ListNode) *ListNode{
    dummyHead := &ListNode{}
    temp, temp1, temp2 := dummyHead, head1, head2
    for temp1 != nil && temp2 != nil {
        if temp1.Val <= temp2.Val {
            temp.Next = temp1
            temp1 = temp1.Next
        } else {
            temp.Next = temp2
            temp2 = temp2.Next
        }
        temp = temp.Next
    }
    if temp1 != nil {
        temp.Next = temp1
    } else if temp2 != nil {
        temp.Next = temp2
    }
    return dummyHead.Next
}
