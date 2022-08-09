package digit

// 数字范围按位与
// RangeBitwiseAnd [m,n]所有数字按位与
func RangeBitwiseAnd(m int, n int) int {
	for m < n {
		n = n & (n - 1)
	}
	return n
}
