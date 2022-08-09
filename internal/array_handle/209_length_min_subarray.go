package arrayhandle

import "math"

// 长度最小的子数组
// 滑动窗口
func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := math.MaxInt
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			res = min(res, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
