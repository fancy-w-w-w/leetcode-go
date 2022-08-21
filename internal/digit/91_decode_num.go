package digit

// 解码方法
// 1~26映射为A~Z，给定一串数字，输出所有可能的字母组合
func NumDecodings(s string) int {
	dp := make([]int, len(s)+1)

	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1]-'0' != 0 {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2]-'0' > 0 && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
