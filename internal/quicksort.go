package internal

import (
	"math/rand"
)

// 快排
func QuickSort(nums []int, start int, end int) []int {
	if start < end {
		partitonIndex := partition(nums, start, end)
		QuickSort(nums, start, partitonIndex-1)
		QuickSort(nums, partitonIndex+1, end)
	}
	return nums
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := left + 1

	for i := index; i <= right; i++ {
		if arr[i] > arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}

	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

// 快排2
// 两路快排
func QuickSortV2(arr []int, l, r int) []int {
	if l >= r {
		return nil
	}

	pivot := partition2(arr, l, r)
	QuickSortV2(arr, l, pivot-1)
	QuickSortV2(arr, pivot+1, r)
	return arr
}

func partition2(arr []int, l, r int) int {
	// 随机选择一个index作为pivot
	pivot := rand.Intn(r-l+1) + l
	arr[l], arr[pivot] = arr[pivot], arr[l]

	v := arr[l]
	i, j := l+1, r
	for {
		// 左边找到第一个不小于v的元素
		for i <= r && arr[i] < v {
			i++
		}
		// 右边找到第一个不大于v的元素
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func swap(a, b *ListNode) {
	a.Val, b.Val = b.Val, a.Val
}

// 单链表快排
func quicksort(head, end *ListNode) {

	if head == end || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	pivot := head.Val
	for fast != nil {
		if fast.Val <= pivot {
			slow = slow.Next
			swap(slow, fast)
		}
		fast = fast.Next
	}
	swap(head, slow)
	quicksort(head, slow)
	quicksort(slow.Next, end)
}
