package matrix

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	res := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	roundTmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > intervals[i-1][1] {
			res = append(res, roundTmp)
			roundTmp = intervals[i]
		} else {
			roundTmp[1] = max(intervals[i][1], roundTmp[1])
			intervals[i] = roundTmp
		}
	}
	res = append(res, roundTmp)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
