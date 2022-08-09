package digit

// TrailingZeroes 给定一个整数 n ，返回 n! 结果中尾随零的数量
// 阶乘后的零
func TrailingZeroes(n int) int {
	res := 0
	for n > 0 {
		n = n / 5
		res += n
	}
	return res
}
