package dynamicresolve

// 最大连续1的个数3
// 给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
// 滑动窗口
func longestOnes(nums []int, k int) (ans int) {
	// lesum为左边0的数量和
	// rsum为当前位置0的数量和
	var left, lsum, rsum int
	for right, v := range nums {
		rsum += 1 - v
		for lsum < rsum-k {
			//
			lsum += 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
