package internal

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
