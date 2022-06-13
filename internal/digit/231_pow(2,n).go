package digit

import "math"

// IsPowerOfTwo 给你一个整数 n，请你判断该整数是否是 2 的幂次方
// 2的n次方
func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// IsPowerOfTwoV2 位运算
func IsPowerOfTwoV2(n int) bool {
	var digit int
	if n == 0 {
		return false
	} else if n > 0 {
		digit = 1
		for digit <= n && digit < math.MaxInt/2 {
			if digit == n {
				return true
			}
			digit <<= 1
		}
	} else {
		for digit >= n && digit >= math.MaxInt/2 {
			if digit == n {
				return true
			}
			digit <<= 1
		}
	}

	return false

}
