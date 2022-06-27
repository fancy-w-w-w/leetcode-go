package arrayhandle

// NextGreaterElements 单调栈
// 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。
// 如果不存在，则输出 -1 。 循环找下一个数
func NextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, 2*n)
	for i := range res {
		res[i] = -1
	}
	nums = append(nums, nums...)
	stack := []int{0}
	for i := 1; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			preElemIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[preElemIndex] = nums[i]
		}
		stack = append(stack, i)
	}
	return res[:n]
}
