package dynamicresolve

// 和等于 k 的最长子数组长度
// nums有正有负
// 给定一个数组 nums 和一个目标值 k，找到和等于 k 的最长连续子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。
// 前缀和+哈希表
func maxSubArrayLen(nums []int, k int) int {
	map1 := make(map[int]int, 0)
	map1[0] = -1
	var preSum int = 0
	res := 0
	for i := range nums {
		preSum += nums[i]
		if _, ok := map1[preSum]; !ok {
			map1[preSum] = i
		}
		if v, ok := map1[preSum-k]; ok {
			res = max(res, i-v)
		}
	}
	return res
}
