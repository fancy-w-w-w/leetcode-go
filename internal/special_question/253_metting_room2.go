package specialquestion

import "sort"

// 会议室2
// 给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，返回 所需会议室的最小数量 。
// 上下车法 结束下车 开始上车
func minMeetingRooms(intervals [][]int) int {
	var nums []int
	// 用个位标识开始和结束
	for _, v := range intervals {
		nums = append(nums, v[0]*10+2)
		nums = append(nums, v[1]*10+1)
	}
	sort.Ints(nums)
	maxRoom := 0
	curNeedRoom := 0
	for _, v := range nums {
		if v%10 == 1 {
			curNeedRoom--
		} else {
			curNeedRoom++
		}
		if curNeedRoom > maxRoom {
			maxRoom = curNeedRoom
		}
	}
	return maxRoom
}
