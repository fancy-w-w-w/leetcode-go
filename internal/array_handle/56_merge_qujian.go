package arrayhandle

import "sort"

// 56. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func MergeIntervals(intervals [][]int) [][]int {
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	tmpInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > tmpInterval[1] {
			res = append(res, tmpInterval)
			tmpInterval = intervals[i]
		} else {
			tmpInterval[1] = max(intervals[i][1], tmpInterval[1])
		}

	}

	res = append(res, tmpInterval)
	return res
}
