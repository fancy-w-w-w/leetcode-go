package arrayhandle

// 字符串乘法
// MultiplyTwoString 字符串相乘
// 1 <= num1.length, num2.length <= 200
func MultiplyTwoString(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	resArr := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			resArr[i+j+1] += x * y
		}

	}
	for i := len(num1) + len(num2) - 1; i > 0; i-- {
		resArr[i-1] += resArr[i] / 10
		resArr[i] %= 10
	}

	index := 0
	// 乘法可能不进位
	// 进位的话结果长度
	if resArr[0] == 0 {
		index++
	}
	jinwei := index
	res := make([]byte, len(num1)+len(num2))
	for ; index < len(num1)+len(num2); index++ {
		res[index] = byte('0' + resArr[index])
	}
	return string(res[jinwei:])
}
