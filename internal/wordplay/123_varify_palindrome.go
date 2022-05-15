package wordplay

import "fmt"

// IsPalindrome 验证回文字符串
func IsPalindrome(s string) bool {
	target := transToString(s)
	if len(target) <= 1 {
		return true
	}
	fmt.Println(target)
	return judgeString(target, len(target)/2, len(target)/2) || judgeString(target, len(target)/2-1, len(target)/2)
}

func judgeString(s string, left int, right int) bool {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			return false
		}
		left--
		right++
	}
	// 是否到达边界
	return left < 0 && right > (len(s)-1)
}

func transToString(s string) string {
	ss := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		str := s[i]
		if ('0' <= str && str <= '9') || ('a' <= str && str <= 'z') {
			ss = append(ss, str)
		} else if 'A' <= str && str <= 'Z' {
			ss = append(ss, byte('a'+str-'A'))
		}
	}
	return string(ss)
}
