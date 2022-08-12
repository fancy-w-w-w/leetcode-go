package arrayhandle

// 缺失的第一个正数
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
func FirstMissingPositive(nums []int) int {
	var res int = len(nums) + 1
	n := len(nums)
	// 负数变为正数
	for index, value := range nums {
		if value <= 0 {
			nums[index] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for index, value := range nums {
		if value > 0 {
			return index + 1
		}
	}
	return res
}
