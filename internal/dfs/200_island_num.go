package dfs

func NumIslands(grid [][]byte) int {
	res := 0
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}

		if grid[i][j] != '1' {
			return
		}

		// 扩散
		grid[i][j] = '0'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 0; i < len(grid); i++ {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
