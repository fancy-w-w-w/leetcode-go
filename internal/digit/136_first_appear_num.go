package digit

// 只出现一次的数字
func SingleNumber(nums []int) int {
	res := 0

	for i := range nums {
		res ^= nums[i]
	}
	return res
}
