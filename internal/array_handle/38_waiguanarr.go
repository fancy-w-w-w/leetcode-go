package arrayhandle

import (
	"strconv"
	"strings"
)

// 外观数列
// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述
func CountAndSay(n int) string {
	firstItem := "1"
	for i := 2; i <= n; i++ {
		oneTurnString := &strings.Builder{}
		for j, start := 0, 0; j < len(firstItem); start = j {
			for j < len(firstItem) && firstItem[j] == firstItem[start] {
				j++
			}
			oneTurnString.WriteString(strconv.Itoa(j - start))
			oneTurnString.WriteByte(firstItem[start])
		}
		firstItem = oneTurnString.String()
	}
	return firstItem
}
