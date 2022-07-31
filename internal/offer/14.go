package offer

import "fmt"

// 剑指 Offer 14- I. 剪绳子
func CuttingRope(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
			fmt.Println(i, dp[i])
		}
	}

	return dp[n]
}
