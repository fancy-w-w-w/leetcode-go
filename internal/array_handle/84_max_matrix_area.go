package arrayhandle

// LargestRectangleArea
// 柱状图中最大的矩形
// 点掉栈
// 最大矩形
func LargestRectangleArea(heights []int) int {
	// stack递增的
	stack := []int{0}
	// 首末填充，防止单调递增或者递减
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	res := 0
	for i := 1; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			topElemIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[topElemIndex]
			w := i - stack[len(stack)-1] - 1
			res = max(res, h*w)
		}
		stack = append(stack, i)
	}
	return res
}
