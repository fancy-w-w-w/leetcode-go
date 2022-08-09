package arrayhandle

import "math"

// FindPeakElement 峰值元素是指其值严格大于左右相邻值的元素。
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
// 寻找波峰
func FindPeakElement(nums []int) int {
	// case 必有波峰
	left, right := 0, len(nums)-1

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > get(mid-1) && nums[mid] > get(mid+1) {
			return mid
		}
		if nums[mid] > get(mid-1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
