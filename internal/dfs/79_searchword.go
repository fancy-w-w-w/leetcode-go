package dfs

type pair struct {
	x, y int
}

var movePath = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 矩阵单词搜索
// ExistWord 矩阵是否存在模式字符串
func ExistWord(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	// 	二维数组标记本轮是否访问过，防止循环dfs
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		// 剪枝
		if board[i][j] != word[k] {
			return false
		}
		// 成功匹配
		if k == len(word)-1 {
			return true
		}

		visited[i][j] = true
		// 回溯消除标记
		defer func() {
			visited[i][j] = false
		}()

		for _, path := range movePath {
			if newi, newj := i+path.x, j+path.y; 0 <= newi && 0 <= newj && newi < h && newj < w && !visited[newi][newj] {
				if check(newi, newj, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
