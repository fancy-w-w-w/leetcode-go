package wordplay

// MinDisplay 给定单词，请两个单词之间最少改变步数 每一步可以change insert delete
func MinDisplay(word1 string, word2 string) int {
	// 空字符串直接返回长度
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	// init dp
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i, _ := range dp {
		dp[i] = make([]int, len2+1)
	}
	// dp[i][0],word1字串到""距离;dp[0][j]""到words2字串距离
	for i := 0; i < len1+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < len2+1; j++ {
		dp[0][j] = j
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	compareStr := func(a, b byte) int {
		if a == b {
			return 0
		}
		return 1
	}
	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {

			dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+compareStr(word1[i-1], word2[j-1])))
		}
	}
	return dp[len1][len2]
}
