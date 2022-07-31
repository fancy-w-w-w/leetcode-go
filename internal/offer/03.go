package offer

// 剑指 Offer 03. 数组中重复的数字
func FindRepeatNumber(nums []int) int {
	map1 := make(map[int]struct{}, 0)

	for i := range nums {
		if _, ok := map1[nums[i]]; ok {
			return nums[i]
		} else {
			map1[nums[i]] = struct{}{}
		}
	}

	return -1
}
