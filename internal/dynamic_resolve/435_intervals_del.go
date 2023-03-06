package dynamicresolve

import "sort"

// 无重叠区间
// 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	if len(intervals) == 0 {
		return 0
	}

	var res int = 1
	rightNum := intervals[0][1]

	for _, p := range intervals[1:] {
		if p[0] < rightNum {
			continue
		}
		res++
		rightNum = p[1]
	}

	return len(intervals) - res
}
