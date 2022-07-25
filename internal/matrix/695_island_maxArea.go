package matrix

// MaxAreaOfIsland 岛屿最大面积
func MaxAreaOfIsland(grid [][]int) int {
	var dfs func(int, int) int
	dfs = func(i1, i2 int) int {
		if i1 < 0 || i2 < 0 || i1 >= len(grid) || i2 >= len(grid[0]) {
			return 0
		}
		if grid[i1][i2] == 0 {
			return 0
		}

		grid[i1][i2] = 0
		return 1 + dfs(i1, i2+1) + dfs(i1+1, i2) + dfs(i1-1, i2) + dfs(i1, i2-1)
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			res = max(res, dfs(i, j))
		}
	}
	return res
}
