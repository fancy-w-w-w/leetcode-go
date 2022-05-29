package wordplay

import (
	"sort"
	"strconv"
	"strings"
)

// LargestNumber 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数
// 数学证明局部相邻两个最优解=整体最优解
func LargestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := strconv.Itoa(x), strconv.Itoa(y)
		return strings.Compare(sx+sy, sy+sx) > 0
	})
	if nums[0] == 0 {
		return "0"
	}
	res := []byte{}
	for _, v := range nums {
		res = append(res, strconv.Itoa(v)...)
	}
	return string(res)
}
