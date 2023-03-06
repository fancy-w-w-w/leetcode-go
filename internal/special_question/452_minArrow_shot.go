package specialquestion

import "sort"

// findMinArrowsShots 用最小的箭引爆气球
// points为pair(中心点，直径最右点)
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	maxRight := points[0][1]

	res := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			res++
		}
	}
	return res
}
