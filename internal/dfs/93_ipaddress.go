package dfs

import (
	"strconv"
	"strings"
)

const SEG_COUNT = 4

var (
	ans      []string
	segments []string
	ipStr    string
)

// RestoreIpAddresses 复原ip地址
// 25525511135->255.255.111.35|2555.255.11.135
func RestoreIpAddresses(s string) []string {
	ans = make([]string, 0)
	segments = make([]string, 5)
	ipStr = s
	dfs(0, 0)
	return ans
}

func dfs(segNum int, segStart int) {
	// 如果找到了四段ip地址并且遍历完字符串，就是一个答案
	if segNum == SEG_COUNT && segStart == len(ipStr) {
		ipAddr := strings.Join(segments, ".")
		ans = append(ans, ipAddr[:len(ipAddr)-1])
	}

	// 没找到四段且到了末尾，返回 || 找到四段没到末尾返回
	if segStart == len(ipStr) || (segNum == SEG_COUNT && segStart < len(ipStr)) {
		return
	}

	// 不能有前导零
	if ipStr[segStart] == '0' {
		segments[segNum] = "0"
		dfs(segNum+1, segStart+1)
	}

	// 普通情况
	singleAddrNum := 0
	for segLenEnd := segStart; segLenEnd < len(ipStr); segLenEnd++ {
		singleAddrNum = singleAddrNum*10 + int(ipStr[segLenEnd]-'0')
		if singleAddrNum == 0 || singleAddrNum > 0xFF {
			break
		}
		segments[segNum] = strconv.FormatInt(int64(singleAddrNum), 10)
		dfs(segNum+1, segLenEnd+1)
	}
}
