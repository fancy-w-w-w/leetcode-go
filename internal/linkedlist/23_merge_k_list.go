package linkedlist

import "container/heap"

type hp struct {
	listNodes []*ListNode
}

func (h *hp) Len() int {
	return len(h.listNodes)
}
func (h *hp) Less(i, j int) bool {
	return h.listNodes[i].Val < h.listNodes[j].Val
}

func (h *hp) Swap(i, j int) {
	h.listNodes[i], h.listNodes[j] = h.listNodes[j], h.listNodes[i]
}

func (h *hp) Push(v interface{}) {
	h.listNodes = append(h.listNodes, v.(*ListNode))
}

func (h *hp) Pop() interface{} {
	old := h.listNodes[h.Len()-1].Val
	h.listNodes = h.listNodes[:h.Len()-1]
	return old
}

// 合并K个升序链表
// 小根堆
func MergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{}
	workNode := dummyHead
	filterList := []*ListNode{}

	// 过滤掉nil空链表
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		filterList = append(filterList, lists[i])
	}
	hp1 := &hp{
		listNodes: filterList,
	}
	heap.Init(hp1)

	for hp1.Len() > 0 {
		nextNode := heap.Pop(hp1).(*ListNode)
		// 链表不为空继续加入堆
		if nextNode != nil && nextNode.Next != nil {
			heap.Push(hp1, nextNode.Next)
		}

		workNode.Next = nextNode
		workNode = workNode.Next
	}

	return dummyHead.Next
}
