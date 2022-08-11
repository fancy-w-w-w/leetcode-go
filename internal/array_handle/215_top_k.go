package arrayhandle

import (
	"container/heap"
	"math/rand"
	"time"
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

// 数组中的第K个最大元素2 快速排序
// 时间复杂度O(n) TOPK
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
