package wordplay

import (
	"strconv"
	"strings"
)

// FindRepeatedDnaSequences 给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)
func FindRepeatedDnaSequences(s string) []string {
	// strA, strC := 1, 2
	// strG, strT := 3, 4
	// for i := 9; i < len(s); i++ {
	// 	tmpStr := s[i-9 : i+1]
	// 	for j := 0; j < 10; j++{
	// 		switch tmpStr[j]:
	// 	}
	// }
	transStrs := strings.ReplaceAll(s, "A", "1")
	transStrs1 := strings.ReplaceAll(transStrs, "C", "2")
	transStrs2 := strings.ReplaceAll(transStrs1, "G", "3")
	targetStr := strings.ReplaceAll(transStrs2, "T", "4")
	res := []string{}

	mapKv := make(map[int64]int, 0)
	for i := 9; i < len(targetStr); i++ {
		tmpStr := targetStr[i-9 : i+1]
		key, _ := strconv.ParseInt(tmpStr, 10, 64)
		if v, ok := mapKv[key]; ok {
			if v == 1 {
				res = append(res, recover(tmpStr))
			}
			mapKv[key]++
		} else {
			mapKv[key] = 1
		}
	}
	return res
}

func recover(s string) string {
	transStrs := strings.ReplaceAll(s, "1", "A")
	transStrs1 := strings.ReplaceAll(transStrs, "2", "C")
	transStrs2 := strings.ReplaceAll(transStrs1, "3", "G")
	targetStr := strings.ReplaceAll(transStrs2, "4", "T")
	return targetStr
}
