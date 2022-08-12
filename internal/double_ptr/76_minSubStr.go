package doubleptr

import (
	"fmt"
	"math"
)

// MinWindow 最小覆盖字串
// 求s中包含t所有字符的最小字串
func MinWindow(s string, t string) string {
	// 待匹配和模式串hash存储字符出现次数
	strT, strRound := make(map[byte]int, 0), make(map[byte]int, 0)

	for i := 0; i < len(t); i++ {
		strT[t[i]]++
	}
	// 检查字串是否包含模式串func
	checkHaveEnoughStr := func() bool {
		for k, v := range strT {
			if strRound[k] < v {
				return false
			}
		}
		return true
	}
	subStrLen := math.MaxInt32
	// 答案左右边界初始化为-1
	ansL, ansR := -1, -1
	for l, r := 0, 0; r < len(s); r++ {
		if r < len(s) && strT[s[r]] > 0 {
			strRound[s[r]]++
		}

		if checkHaveEnoughStr() && l <= r {

			// 字串长度更短
			if r-l+1 < subStrLen {
				subStrLen = r - l + 1
				ansL, ansR = l, l+subStrLen
				fmt.Println(s[ansL:ansR])
			}
			// 滑动窗口
			if _, ok := strT[s[l]]; ok {
				strRound[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]

}

// 最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func MinWindowV2(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
