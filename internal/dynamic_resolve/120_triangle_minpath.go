package dynamicresolve

import "math"

// MinimumTatalPath 三角形最小路径和
func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := [2][]int{}
	for i := 0; i < 2; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		curr := i % 2
		prev := 1 - curr
		f[curr][0] = f[prev][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[curr][j] = min(f[prev][j-1], f[prev][j]) + triangle[i][j]
		}
		f[curr][i] = f[prev][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[(n-1)%2][i])
	}
	return ans
}
