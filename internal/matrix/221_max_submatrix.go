package matrix

// MaximalSquare 求最大的正方形
//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积
func MaximalSquare(matrix [][]byte) int {
	if len(matrix) <= 1 || len(matrix[0]) <= 1 {
		for i := range matrix {
			for j := range matrix[i] {
				if matrix[i][j] == '1' {
					return 1
				}
			}
		}
		return 0
	}

	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	res := 0

	for i := range matrix {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			res = 1
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			res = 1
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			res = max(res, dp[i][j])
		}
	}
	return res * res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
