package arrayhandle

// FindAnagrams
//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 定长滑动窗口
func FindAnagrams(s string, p string) (ans []int) {
	m, n := len(s), len(p)
	if m < n {
		return ans
	}
	var cntsP, cnts [26]int
	for i := 0; i < n; i++ {
		cntsP[p[i]-byte('a')]++
	}
	for i := 0; i < m; i++ {
		cnts[s[i]-byte('a')]++
		if i >= n-1 {
			if cnts == cntsP {
				ans = append(ans, i-n+1)
			}
			cnts[s[i-n+1]-byte('a')]--
		}
	}
	return ans
}
