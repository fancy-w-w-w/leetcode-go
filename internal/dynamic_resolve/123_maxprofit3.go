package dynamicresolve

func MaxProfitV3x(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][0], dp[0][2], dp[0][4] = 0, 0, 0
	dp[0][1], dp[0][3] = -prices[0], -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	// 返回的情况需要注意，不需比较
	return dp[len(prices)-1][4]
}
