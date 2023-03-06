package specialquestion

import "sort"

// 455. 分发饼干
//
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	indexg := 0
	indexs := 0
	res := 0
	for indexs < len(s) && indexg < len(g) {

		if g[indexg] <= s[indexs] {
			res++
			indexg++
		}
		indexs++
	}
	return res
}
