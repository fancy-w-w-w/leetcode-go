package dynamicresolve

// 串联所有单词的子串
// 给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 滑动窗口
func FindSubstring(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n <= ls; i++ {
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}
