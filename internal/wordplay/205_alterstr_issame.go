package wordplay

// IsIsomorphic 两字符串是否是简单映射
// 需要两个hash
func IsIsomorphic(s string, t string) bool {
	alterMap1 := make(map[byte]byte, 0)
	if len(s) != len(t) {
		return false
	}
	alterMap2 := make(map[byte]byte, 0)

	for i, j := 0, 0; i < len(s) && j < len(t); i, j = i+1, j+1 {
		if str, ok := alterMap1[t[j]]; ok && str != s[i] {
			return false
		}

		if str, ok := alterMap2[s[i]]; ok && str != t[j] {
			return false
		}
		alterMap2[s[i]] = t[j]

		alterMap1[t[j]] = s[i]
	}
	return true
}
