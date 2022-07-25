package matrix

func orangesRotting(grid [][]int) int {
	cnts, cost, m, n, queue := 0, 0, len(grid), len(grid[0]), [][]int{}
	for i, row := range grid {
		for j, v := range row {
			if v > 0 {
				cnts++
				if v == 2 {
					queue = append(queue, []int{i, j})
				}
			}
		}
	}
	for len(queue) > 0 {
		l := len(queue)
		cnts -= l
		for i := 0; i < l; i++ {
			point := queue[0]
			queue = queue[1:]
			for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				nx, ny := point[0]+d[0], point[1]+d[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					queue = append(queue, []int{nx, ny})
				}
			}
		}
		if len(queue) > 0 {
			cost++
		}
	}
	if cnts == 0 {
		return cost
	}
	return -1
}
