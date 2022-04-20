package dynamicresolve

func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for index, _ := range dp {
		dp[index] = make([]int, n)
	}

	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dp[i][j] += dp[i][j-1]
			} else if j == 0 {
				dp[i][j] += dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func UniquePathsV2(obstacleGrid [][]int) int {
	rowLen := len(obstacleGrid)
	lineLen := len(obstacleGrid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if obstacleGrid[i][j] == 1 {
			return 0
		}
		if i == rowLen-1 && j == lineLen-1 {
			return 1
		}

		if i == rowLen-1 {
			return dfs(i, j+1)
		} else if j == lineLen-1 {
			return dfs(i+1, j)
		}
		return dfs(i+1, j) + dfs(i, j+1)
	}

	return dfs(0, 0)
}

// 滚动数组优化dp空间（累加）
func UniquePathsWithObstaclesV3(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}
