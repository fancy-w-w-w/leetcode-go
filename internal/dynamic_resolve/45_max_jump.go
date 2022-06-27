package dynamicresolve

// Jump  跳跃游戏 II 最短步数
func Jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	// 因为总能跳到最后一个位置，所以不算最后一个数组元素
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
