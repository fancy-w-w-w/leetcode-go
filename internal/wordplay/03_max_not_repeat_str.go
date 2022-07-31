package wordplay

// LengthOfLongestSubstring 无重复字符的最长子串
// 滑动窗口
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	map1 := make(map[byte]int, 0)
	right, res := 0, 1
	for i := 0; i < len(s); i++ {
		delete(map1, s[i])

		for right < len(s) && map1[s[right]] == 0 {
			map1[s[right]]++
			right++
		}
		res = max(res, right-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
