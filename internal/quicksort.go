package internal

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
