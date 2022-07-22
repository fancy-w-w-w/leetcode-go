package arrayhandle

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
// 滑动窗口
// 只能解决没有负数的array
func SubarraySum(nums []int, k int) int {
	// window := make([]int, 0)
	left, right := 0, 0
	target := 0
	res := 0
	for right <= len(nums)-1 {
		// window = append(window, nums[right])
		target += nums[right]
		for left <= right && target >= k {
			if target == k {
				res++
			}
			// window = window[:len(window)-1]
			target -= nums[left]
			left++
		}
	}
	return res
}

// SubarraySum2 负数数组
// 前缀和+哈希表解法
func SubarraySum2(nums []int, k int) int {
	preSum := make(map[int]int, 0)
	res, pre := 0, 0
	// 需要初始化前缀和为0的情况
	preSum[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if count, ok := preSum[pre-k]; ok {
			res += count
		}
		preSum[pre] += 1
	}
	return res
}
