package binarysearch

// FindMinRotated 寻找旋转过后的数组的最小值 二分法 数组元素可重复
func FindMinRotated(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else if nums[pivot] > nums[high] {
			low = pivot + 1
		} else {
			// 重复 high--
			high--
		}
	}
	return nums[low]
}
