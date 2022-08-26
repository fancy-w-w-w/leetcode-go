package arrayhandle

// 在排序数组中查找元素的第一个和最后一个位置
// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置
func searchRange(nums []int, target int) []int {
	leftRes, rightRes := -1, -1
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			if mid == 0 || (mid > 0 && nums[mid-1] != target) {
				leftRes = mid
				break
			}
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	l, r = 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			if mid == n-1 || (mid < n-1 && nums[mid+1] != target) {
				rightRes = mid
			}
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{leftRes, rightRes}
}
