package dfs

// CombinationSum3
// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
// 只使用数字1到9
// 每个数字 最多使用一次
func CombinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var dfs func([]int, int, int)
	dfs = func(combine []int, sum int, num int) {
		if sum == n {
			if len(combine) == k {
				res = append(res, append([]int{}, combine...))
			}
		} else if sum > n {
			return
		}
		// len > k
		if len(combine) > k {
			return
		}

		for i := num + 1; i <= 9; i++ {
			combine = append(combine, i)
			dfs(combine, sum+i, i)
			combine = combine[:len(combine)-1]
		}

	}

	dfs([]int{}, 0, 0)
	return res
}
