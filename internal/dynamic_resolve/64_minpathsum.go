package dynamicresolve

import "fmt"

// MinPathSum 最小路径和
func MinPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j == 0 {
				dp[j] += grid[i][j]
			} else if i == 0 {
				dp[j] = grid[i][j] + dp[j-1]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
			}
		}
		fmt.Println(dp)
	}

	return dp[len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
