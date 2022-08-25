package arrayhandle

// 寻找重复数
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数
// 快慢指针终将相遇解法
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}

	// 求环的入口
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
