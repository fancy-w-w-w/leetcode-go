package dynamicresolve

// 最多交易k次
func MaxProfitV4(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}

	// 第1,3...天持有股票
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			// 持有
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			// 不持有
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}
