package arrayhandle

import "fmt"

func Trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	stack := []int{0}
	res := 0

	for i := 1; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			low := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				h := min(height[stack[len(stack)-1]], height[i]) - height[low]
				w := i - stack[len(stack)-1] - 1
				res += h * w
			}
		}
		stack = append(stack, i)
		fmt.Println(stack, res)
	}
	return res

}

// stack 单调递减
