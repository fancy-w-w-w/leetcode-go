package digit

// CountPrimes 统计质数个数，采用标记筛选法
// 计数质数
func CountPrimes(n int) int {
	// 筛选
	isPrimes := make([]bool, n)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	res := 0
	for i := 2; i < n; i++ {
		if !isPrimes[i] {
			continue
		}
		res++
		for j := i * 2; j < n; j += i {
			isPrimes[j] = false
		}
	}
	return res
}
