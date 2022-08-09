package arrayhandle

import "fmt"

// FindMissingRanges 给定一个排序的整数数组 nums ，其中元素的范围在 闭区间 [lower, upper] 当中，返回不包含在数组中的缺失区间。
// 缺失的区间
func FindMissingRanges(nums []int, lower int, upper int) []string {
	preStart := lower
	res := []string{}
	// 重点
	nums = append(nums, upper+1)
	for i := 0; i < len(nums); i++ {
		if nums[i] < lower {
			continue
		}
		if nums[i] == preStart {
			preStart++
			continue
		}
		if nums[i] == preStart+1 {
			res = append(res, fmt.Sprintf("%d", preStart))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", preStart, nums[i]-1))
		}
		preStart = nums[i] + 1
	}
	return res
}
