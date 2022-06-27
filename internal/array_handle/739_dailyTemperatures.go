package arrayhandle

// DailyTemperatures
//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指在第 i 天之后，才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 单调栈
func DailyTemperatures(temperatures []int) []int {
	stack := []int{0}
	res := make([]int, len(temperatures))
	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return res
}
