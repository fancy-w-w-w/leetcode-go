package offer

// 剑指 Offer 13. 机器人的运动范围
func MovingCount(m int, n int, k int) int {
	res := 0
	visted := make([][]bool, m)
	for i := range visted {
		visted[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i int, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if visted[i][j] {
			return
		}
		sumNum := i%10 + i/10 + j/10 + j%10 + i/100 + j/100

		visted[i][j] = true
		if sumNum <= k {
			res++
		} else {
			return
		}
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i+1, j)
	}
	dfs(0, 0)
	return res
}
