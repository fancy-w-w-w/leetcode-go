package arrayhandle

// 去除重复字母
// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
// 贪心 单调栈
func RemoveDuplicateLetters(s string) string {
	// 记录字符剩余的次数
	leftArr := [26]int{}
	stack := []byte{}

	for i := range s {
		leftArr[s[i]-'a']++
	}

	instack := [26]bool{}
	for i := range s {
		ch := s[i]
		// 已经在栈中了，说明重复了，直接跳过
		if instack[ch-'a'] {
			leftArr[ch-'a']--
			continue
		}

		//不在栈中
		// 单调栈
		for len(stack) > 0 && ch < stack[len(stack)-1] {
			lastCh := stack[len(stack)-1] - 'a'
			// 如果剩余字符中没有了 不能出栈
			if leftArr[lastCh] == 0 {
				break
			}
			stack = stack[:len(stack)-1]
			instack[lastCh] = false
		}

		stack = append(stack, ch)
		instack[s[i]-'a'] = true
		leftArr[ch-'a']--
	}

	return string(stack)
}
