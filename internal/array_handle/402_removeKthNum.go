package arrayhandle

import "strings"

// 移掉K位数字
// 给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。
// 贪心+单调栈
func RemoveKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	stack := []byte{}
	for i := range num {
		digit := num[i]
		// 单调栈
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			// 删除元素
			stack = stack[:len(stack)-1]
			k--
		}
	}

	// 如果k没减完，单调栈剩下的元素全部单调，因此直接去掉尾部数字
	stack = stack[:len(stack)-k]

	// 去掉前导零
	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		return "0"
	}

	return res
}
