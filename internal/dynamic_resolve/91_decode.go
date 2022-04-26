package dynamicresolve

// numDecoding 求字符映射成num数组过后，复原有多少种串
func NumDecoding(s string) int {
	dp := make([]int, len(s)+1)

	dp[0] = 1

	// 上界为字符串末尾后一位
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
