package wordplay

import (
	"strconv"
	"strings"
)

func CompareVersion(version1 string, version2 string) int {
	versionAStrs := strings.Split(version1, ".")
	versionBStrs := strings.Split(version2, ".")

	var i, j int
	for i, j = 0, 0; i < len(versionAStrs) && j < len(versionBStrs); i, j = i+1, j+1 {
		switch compareSubStr(versionAStrs[i], versionBStrs[j]) {
		case 1:
			return 1
		case -1:
			return -1
		case 0:
			continue
		}
	}
	for i < len(versionAStrs) {
		num, _ := strconv.ParseInt(versionAStrs[i], 10, 32)
		if num == 0 {
			i++
			continue
		}
		if num > 0 {
			return 1
		}
		return -1
	}
	for j < len(versionBStrs) {
		num, _ := strconv.ParseInt(versionBStrs[j], 10, 32)
		if num == 0 {
			j++
			continue
		}
		if num > 0 {
			return -1
		}
		return 1
	}
	return 0
}

func compareSubStr(a, b string) int {
	num1, _ := strconv.ParseInt(a, 10, 32)
	num2, _ := strconv.ParseInt(b, 10, 32)
	if num1 == num2 {
		return 0
	}
	if num1 > num2 {
		return 1
	}
	return -1
}
