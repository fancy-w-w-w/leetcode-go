package digit

// HammingWeight 输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）
func HammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if ((1 << i) & num) > 0 {
			res++
		}

	}
	return res
}
