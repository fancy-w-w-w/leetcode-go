package binarysearch

// 无重复情况
func SearchNum(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[n-1] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// 有重复情况
/*
对于数组中有重复元素的情况，二分查找时可能会有 a[l]=a[\textit{mid}]=a[r]a[l]=a[mid]=a[r]，此时无法判断区间 [l,\textit{mid}][l,mid] 和区间 [\textit{mid}+1,r][mid+1,r] 哪个是有序的。

例如 \textit{nums}=[3,1,2,3,3,3,3]nums=[3,1,2,3,3,3,3]，\textit{target}=2target=2，首次二分时无法判断区间 [0,3][0,3] 和区间 [4,6][4,6] 哪个是有序的。

对于这种情况，我们只能将当前二分区间的左边界加一，右边界减一，然后在新区间上继续二分查找。
*/
func SearchNumV2(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		// 重复元素跳过
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
			continue
		} else if nums[l] <= nums[mid] {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[l] > nums[mid] {
			if nums[mid] < target && nums[n-1] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
