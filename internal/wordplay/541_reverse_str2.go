package wordplay

/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func ReverseStr(s string, k int) string {
	byteStr := []byte(s)
	n := len(byteStr)
	i := 0
	for ; i+2*k < n; i += 2 * k {
		reverseSubStr(byteStr, i, i+k-1)
	}
	if i+k < n {
		reverseSubStr(byteStr, i, i+k-1)
	} else {
		reverseSubStr(byteStr, i, n-1)

	}
	return string(byteStr)
}

func reverseSubStr(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// @lc code=end
