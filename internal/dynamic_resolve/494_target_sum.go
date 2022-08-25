package dynamicresolve

// 目标和
// 问题转化成在数组 \textit{nums}nums 中选取若干元素，使得这些元素之和等于 \textit{neg}neg，计算选取元素的方案数
// 记数组的元素和为 \textit{sum}sum，添加 \texttt{-}- 号的元素之和为 \textit{neg}neg
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	n, neg := len(nums), diff/2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[n][neg]
}
