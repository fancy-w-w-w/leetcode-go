package arrayhandle

import "sort"

// 根据身高重建队列
// people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人
// 当我们放入第 ii 个人时，只需要将其插入队列中，使得他的前面恰好有ki个人就好了
func ReconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})

	for i := range people {
		res = append(res[:people[i][1]], append([][]int{people[i]}, res[people[i][1]:]...)...)
	}
	return res
}
