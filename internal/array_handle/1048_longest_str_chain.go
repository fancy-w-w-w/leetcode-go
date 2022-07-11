package arrayhandle

import "sort"

func LongestStrChain(words []string) int {
	res := 0
	dp := make([]int, len(words))
	for i := range dp {
		dp[i] = 1
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i := range words {
		for j := 0; j < i; j++ {
			if len(words[i])-len(words[j]) != 1 {
				continue
			}
			if isSubStr(words[i], words[j]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	for i := range dp {
		res = max(res, dp[i])
	}
	return res
}

func isSubStr(s, t string) bool {

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] != t[j] {
			i++
			continue
		}
		i++
		j++
	}
	return j == len(t)
}
