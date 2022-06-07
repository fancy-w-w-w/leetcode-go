package digit

// IsHappy 快乐数
func IsHappy(n int) bool {
	getSquareSum := func(m int) int {
		sum := 0
		for m != 0 {
			sum += (m % 10) * (m % 10)
			m /= 10
		}
		return sum
	}
	// 快慢指针，有环最终会相遇
	fast, slow := n, n
	slow = getSquareSum(slow)
	fast = getSquareSum(fast)
	fast = getSquareSum(fast)
	for fast != slow {
		slow = getSquareSum(slow)
		fast = getSquareSum(fast)
		fast = getSquareSum(fast)
	}
	if fast == 1 {
		return true
	}
	return false
}
