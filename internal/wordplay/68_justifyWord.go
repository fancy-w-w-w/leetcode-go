package wordplay

import "strings"

// FullJustify 填充单词
// 每一行最大化存储单词 两个单词之间必须有空格 空格数量应该均匀，且左边大于右边
// 最后一行单词之间有一个空格 单单词左对齐 填充空格
func FullJustify(words []string, maxWidth int) []string {
	right, n := 0, len(words)
	ans := make([]string, 0)
	// 遍历每一行
	for {
		// 第一个单词位置
		left := right
		// 统计这一行单词之和
		sumLen := 0

		// 循环确认当前行可以放多少个单词，单词之间有空格 right-left=多少个空格
		for right < n && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		// 如果当前行只有一个单词，该单词左对齐， 行末填充空格
		if right == n {
			s := strings.Join(words[left:], " ")
			ans = append(ans, s+getBlankStr(maxWidth-len(s)))
			return ans
		}

		numWords := right - left
		numSpaces := maxWidth - sumLen

		// 当前行只有一个单词
		if numWords == 1 {
			ans = append(ans, words[left]+getBlankStr(numSpaces))
			continue
		}

		// 当前行不止一个单词
		// 计算平均的插入的空格数
		avgSpaces := numSpaces / (numWords - 1)
		extraSpaces := numSpaces % (numWords - 1)
		// 前一个+1是slice 后一个+1是多个空格
		s1 := strings.Join(words[left:left+extraSpaces+1], getBlankStr(avgSpaces+1))
		s2 := strings.Join(words[left+extraSpaces+1:right], getBlankStr(avgSpaces))

		ans = append(ans, s1+getBlankStr(avgSpaces)+s2)

	}

}

// getBlankStr获得n个空格的string
func getBlankStr(n int) string {
	return strings.Repeat(" ", n)
}
