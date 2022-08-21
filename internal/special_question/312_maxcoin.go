package specialquestion

// 戳气球
// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
// 求所能获得硬币的最大数量。
// 记忆化搜索
func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	val[0], val[n+1] = 1, 1
	rec := make([][]int, n+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, n+2)
		for j := 0; j < len(rec[i]); j++ {
			rec[i][j] = -1
		}
	}
	return solve(0, n+1, val, rec)
}

func solve(left, right int, val []int, rec [][]int) int {
	if left >= right-1 {
		return 0
	}
	if rec[left][right] != -1 {
		return rec[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := val[left] * val[i] * val[right]
		// 左边的max+右边的max
		sum += solve(left, i, val, rec) + solve(i, right, val, rec)
		rec[left][right] = max(rec[left][right], sum)
	}
	return rec[left][right]
}
