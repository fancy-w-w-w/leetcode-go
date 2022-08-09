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

// 最小路径和
func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}
	dp := make([][]int, len(grid))
	for k, v := range grid {
		dp[k] = make([]int, len(v)) //初始化dp数组
	}
	dp[0][0] = grid[0][0] //动态规划边界
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j] //动态规划递推关系式
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j] //动态规划递推关系式
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j] ////动态规划递推关系式
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
