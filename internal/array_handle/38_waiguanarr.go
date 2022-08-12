package arrayhandle

import (
	"strconv"
	"strings"
)

// 外观数列
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
