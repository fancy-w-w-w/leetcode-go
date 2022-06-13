package wordplay

import "strconv"

// DiffWaysToCompute 不同顺序求值
// 输入：expression = "2-1-1"
// 输出：[0,2]
// 解释：
// ((2-1)-1) = 0
// (2-(1-1)) = 2
func DiffWaysToCompute(input string) []int {
	// 如果是数字，直接返回
	if digit, ok := isDigit(input); ok {
		return []int{digit}
	}

	// 空切片
	var res []int
	// 遍历字符串
	for index, c := range input {
		tmpC := string(c)
		if tmpC != "+" && tmpC != "-" && tmpC != "*" {
			continue
		}
		// 如果是运算符，则计算左右两边的算式
		left := DiffWaysToCompute(input[:index])
		right := DiffWaysToCompute(input[index+1:])
		for _, leftNum := range left {
			for _, rightNum := range right {
				var addNum int
				if tmpC == "+" {
					addNum = leftNum + rightNum
				} else if tmpC == "-" {
					addNum = leftNum - rightNum
				} else {
					addNum = leftNum * rightNum
				}
				res = append(res, addNum)
			}
		}

	}

	return res

}

// 判断是否为全数字
func isDigit(input string) (int, bool) {
	var digit int
	digit, err := strconv.Atoi(input)
	if err != nil {
		return 0, false
	}
	return digit, true
}
