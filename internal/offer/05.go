package offer

// 剑指 Offer 05. 替换空格
// 把字符串 s 中的每个空格替换成"%20"
func ReplaceSpace(s string) string {
	res := []byte{}

	for i := range s {
		if s[i] == ' ' {
			res = append(res, []byte("%20")...)
			continue
		}
		res = append(res, s[i])
	}

	return string(res)
}
