package wordplay

import "math"

// ShortestDistance
// 给定一个字符串数组 wordDict 和两个已经存在于该数组中的不同的字符串 word1 和 word2 。返回列表中这两个单词之间的最短距离。
func ShortestDistance(wordsDict []string, word1 string, word2 string) int {
	lastIndex1, lastIndex2 := -1, -1
	res := math.MaxInt
	for i := range wordsDict {
		if wordsDict[i] == word1 {
			if lastIndex2 != -1 {
				res = min(res, i-lastIndex2)
			}
			lastIndex1 = i
		} else if wordsDict[i] == word2 {
			if lastIndex1 != -1 {
				res = min(res, i-lastIndex1)
			}
			lastIndex2 = i
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
