package digit

import (
	"strconv"

	"github.com/emirpasic/gods/lists/arraylist"
)

// GetPermutation 全排列，数字逻辑题
func GetPermutation(n int, k int) string {
	// 题目转化为全排列求从小到大下标第k-1的排列
	k--

	// 阶乘
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	res := make([]byte, 0)

	list := arraylist.New()

	for i := 1; i <= n; i++ {
		list.Add(i)
	}

	for i := n - 1; i >= 0; i-- {
		index := k / factorial[i]
		tmpValue, _ := list.Get(index)
		tmpSaveStr := strconv.Itoa(tmpValue.(int))
		res = append(res, []byte(tmpSaveStr)...)
		list.Remove(index)
		k -= index * factorial[i]
	}

	return string(res)
}
