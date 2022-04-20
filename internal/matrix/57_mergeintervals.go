package matrix

import "fmt"

func InsertIntervalsV2(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	isInsert := false
	for index, value := range intervals {
		if value[0] >= newInterval[0] {
			rear := append([][]int{}, intervals[index:]...)
			intervals = append(append(intervals[:index], newInterval), rear...)
			isInsert = true
		}
	}

	if !isInsert {
		intervals = append(intervals, newInterval)
	}
	fmt.Println(intervals)

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
