package dfs

// 删除有效括号
//给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
// 返回所有可能的结果。答案可以按 任意顺序 返回。
func removeInvalidParentheses(s string) []string {
	l, r := 0, 0
	res := []string{}
	for _, v := range s {
		if v == '(' {
			l++
		} else if v == ')' {
			if l == 0 {
				r++
			} else {
				l--
			}
		}
	}
	dfs1(s, 0, l, r, &res)
	return res

}

func dfs1(s string, start, l, r int, res *[]string) {
	if l == 0 && r == 0 {
		if isValid(s) {
			*res = append(*res, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] {
			continue
		}

		if s[i] == ')' && r > 0 {
			cur := s
			cur = cur[:i] + cur[i+1:]
			dfs1(cur, i, l, r-1, res)
		} else if s[i] == '(' && l > 0 {
			cur := s
			cur = cur[:i] + cur[i+1:]
			dfs1(cur, i, l-1, r, res)
		}
	}

}

func isValid(s string) bool {
	cnt := 0
	for _, v := range s {
		if v == '(' {
			cnt++
		}
		if v == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}
