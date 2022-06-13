package arrayhandle

import "sort"

// CanAttendMeetings
func CanAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := range intervals {
		if i == 0 {
			continue
		}
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
