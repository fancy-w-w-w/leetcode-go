package arrayhandle

import (
	"container/heap"
)

// heap一般用数组
type Heap struct {
	array []int
}

func (h *Heap) Len() int {
	return len(h.array)
}

func (h *Heap) Less(i, j int) bool {
	return h.array[i] < h.array[j]
}

func (h *Heap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *Heap) Push(x any) {
	h.array = append(h.array, x.(int))
}

func (h *Heap) Pop() any {
	res := h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	return res
}

// type Heap []int

// func(h *Heap) Len() int{
// 	return len(*h)
// }

// func(h *Heap) Less(i, j int) bool{
// 	return (*h)[i] < (*h)[j]
// }

// func(h *Heap) Swap(i, j int){
// 	(*h)[i] , (*h)[j] = (*h)[j] , (*h)[i]
// }

// func(h *Heap) Push(x any){
// 	*h = append(*h, x.(int))
// }

// func(h *Heap) Pop() any{
// 	res := (*h)[h.Len()-1]
// 	(*h) = (*h)[:h.Len()-1]
// 	return res
// }

// FindKthLargest 数组中的第K个最大元素
func FindKthLargest(nums []int, k int) int {
	heap1 := &Heap{nums}
	heap.Init(heap1)
	for k > 1 {
		k--
		heap.Pop(heap1)
	}
	return heap.Pop(heap1).(int)
}
