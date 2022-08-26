package digit

import "math"

// 整数反转
// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
