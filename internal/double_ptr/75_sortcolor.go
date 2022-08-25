package doubleptr

// 颜色分类
// SortColors 相同颜色（0,1,2）的元素从小到大排序，原地算法
func sortColors(nums []int) {
	// p0 p1 分别指向0和1连续的下一个位置
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			// 防止i和p0互换后p0,nums[i]为1
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
