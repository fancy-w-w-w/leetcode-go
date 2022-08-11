package dynamicresolve

import "math"

// 零钱兑换
// 返回硬币最少个数
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, coin := range coins {
		for j := 0; j <= amount; j++ {
			if j >= coin {
				dp[j] = min(dp[j], dp[j-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 零钱兑换
// 返回硬币组合数
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
