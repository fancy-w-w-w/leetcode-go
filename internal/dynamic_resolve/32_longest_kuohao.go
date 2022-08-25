package dynamicresolve

// LongestValidParentheses 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
// 动态规划
func LongestValidParentheses(s string) int {
	dp := make([]int, len(s))
	var res int
	for i := 1; i < len(s); i++ {
		tmpStr := s[i]
		if tmpStr == ')' {
			// 1
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				} // 2前一个字符的最长长度之前字符为“（”，说明匹配了
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				// 如果前面还能匹配，继续累加
				// -2 是因为-1的基础上再减去'('
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			res = maxNum(res, dp[i])
		}
	}
	return res
}

func maxNum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
