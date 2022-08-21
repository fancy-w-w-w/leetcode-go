package arrayhandle

import "math"

// 最短无序连续子数组
// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
// 请你找出符合题意的 最短 子数组，并输出它的长度。

func FindUnsortedSubarray(nums []int) int {
	// 利用升序特性
	// 从左到右遍历, 记录当前遍历数的最大值, 最后一个小于最大值的即 需要倒置数组的右边索引
	// 从右到左遍历, 记录当前遍历数的最小值, 最后一个大于最小值的即 需要倒置数组的左边索引
	// 一次遍历
	n := len(nums)
	minNum, maxNum := math.MaxInt32, math.MinInt32
	left, right := -1, -1
	for i := range nums {
		if maxNum > nums[i] {
			// 更新右边索引
			right = i
		} else {
			// 更新最大值
			maxNum = nums[i]
		}
		// 倒序遍历
		if minNum < nums[n-1-i] {
			left = n - 1 - i
		} else {
			minNum = nums[n-1-i]
		}
	}

	// 排除
	if right == -1 || left == -1 {
		return 0
	}
	return right - left + 1
}
