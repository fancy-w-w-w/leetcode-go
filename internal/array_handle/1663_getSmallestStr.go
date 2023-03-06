package arrayhandle

// 1663. 具有给定数值的最小字符串
// 贪心
func getSmallestString(n int, k int) string {
	if k > 26*n || k < n {
		return ""
	}

	res := make([]byte, n)
	for i := range res {
		res[i] = 'a'
	}

	k -= n
	// 先填充a
	for i := n - 1; i >= 0; i-- {
		if k <= 25 {
			res[i] = byte(int(res[i]) + k)
			break
		} else {
			k -= 25
			res[i] = 'z'
		}
	}

	return string(res)
}
