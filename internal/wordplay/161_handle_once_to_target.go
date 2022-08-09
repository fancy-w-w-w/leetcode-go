package wordplay

// IsOneEditDistance给定两个字符串 s 和 t ，如果它们的编辑距离为 1 ，则返回 true ，否则返回 false 。
// 字符串 s 和字符串 t 之间满足编辑距离等于 1 有三种可能的情形：
// 往 s 中插入 恰好一个 字符得到 t
// 从 s 中删除 恰好一个 字符得到 t
// 在 s 中用 一个不同的字符 替换 恰好一个 字符得到 t
// 相隔为1的编辑距离
func IsOneEditDistance(s string, t string) bool {
	return isAlter([]byte(s), t) || isInsert([]byte(s), t) || isDel([]byte(s), t)
}

func isInsert(s []byte, t string) bool {
	if len(s) != len(t)-1 {
		return false
	}
	var isDiff bool
	var diffIndex int
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			isDiff = true
			diffIndex = i
			break
		}
	}
	if !isDiff {
		return true
	}

	// 切片插入元素有坑，需要注意
	return string(s[:diffIndex])+string(t[diffIndex])+string(s[diffIndex:]) == t
}

func isDel(s []byte, t string) bool {
	if len(s) != len(t)+1 {
		return false
	}
	var isDiff bool
	var diffIndex int
	for i := 0; i < len(t); i++ {
		if s[i] != t[i] {
			isDiff = true
			diffIndex = i
			break
		}
	}
	if !isDiff {
		return true
	}
	return string(append(s[:diffIndex], s[diffIndex+1:]...)) == t
}

func isAlter(s []byte, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var isDiff bool
	var diffIndex int
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			isDiff = true
			diffIndex = i
			break
		}
	}
	if !isDiff {
		return false
	}
	return string(append(append(s[:diffIndex], t[diffIndex]), s[diffIndex+1:]...)) == t
}
