package dynamicresolve

// 买卖股票的最佳时机 买入和售出一次
func MaxProfitV1(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][1] = -prices[0]
	dp[0][0] = 0
	for i := 1; i < length; i++ {
		dp[i][1] = max(dp[i-1][1], -prices[i])
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	}
	return dp[length-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
