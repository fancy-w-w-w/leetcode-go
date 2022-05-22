package doubleptr

// LengthOfLongestSubstringTwoDistinct 159. 至多包含两个不同字符的最长子串
func LengthOfLongestSubstringTwoDistinct(s string) int {
	left, right := 0, 0
	res := 0
	storeElem := make(map[byte]int, 0)
	for right < len(s) {
		storeElem[s[right]]++
		right++
		for len(storeElem) > 2 {
			storeElem[s[left]]--
			if storeElem[s[left]] == 0 {
				delete(storeElem, s[left])
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
