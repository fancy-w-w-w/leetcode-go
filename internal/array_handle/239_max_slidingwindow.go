package arrayhandle

import (
	"container/heap"
	"fmt"
	"sort"
)

// 最大滑动窗口
func MaxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		// 头->尾单调递减队列，如果加入元素大于尾元素，需要一直往前找，尾部元素出队
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	// 初始化单调队列
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		// 调整单调队列
		push(i)

		// 滑动窗口移动
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
		fmt.Println(q)
	}
	return ans

}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func MaxSlidingWindow2(nums []int, k int) []int {
	q := &hp{nums}
	heap.Init(q)
	for len(q.IntSlice) != 0 {
		fmt.Println(heap.Pop(q))
	}
	return q.IntSlice

}
