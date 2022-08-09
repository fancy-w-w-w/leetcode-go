package arrayhandle

import "strconv"

// GetPermutation 排列序列
// 给定 n 和 k，返回第 k 个排列。
func GetPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	// n!的值数组
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	// k需要减一
	k--

	ans := ""
	valid := make([]int, n+1)

	// 选择数字 1代表没被选中
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		// 第i个位置的首元素
		order := k/factorial[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		// 求下一位
		k %= factorial[n-i]
	}
	return ans
}
