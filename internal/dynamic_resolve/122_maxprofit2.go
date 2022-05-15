package dynamicresolve

func MaxProfitV2(prices []int) int {
	dp := make([][]int, len(prices))
	length := len(prices)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 唯一一处差别，可以多次买卖

	}
	return dp[length-1][0]
}
