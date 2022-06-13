package wordplay

// IsAnagram 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
func IsAnagram(s string, t string) bool {
	originStr := [26]int{}
	dStr := [26]int{}

	for i := range s {
		originStr[s[i]-'a']++
	}
	for j := range t {
		dStr[t[j]-'a']++
	}

	for i := 0; i < 26; i++ {
		if originStr[i] != dStr[i] {
			return false
		}
	}
	return true
}
