package binarysearch

// 寻找旋转排序数组中的最小值
// 不重复
func FindMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		midValue := nums[mid]
		if midValue > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

// FindMinRotated 寻找旋转过后的数组的最小值 二分法 数组元素可重复
// 寻找旋转排序数组中的最小值2
// 可能重复
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
