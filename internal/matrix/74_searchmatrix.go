package matrix

// SearchMatrix 快速搜索有序二维数组找到target
func SearchMatrix(matrix [][]int, target int) bool {
	row, line := len(matrix)-1, len(matrix[0])-1

	for row >= 0 && line >= 0 {
		if line < 0 || row < 0 {
			return false
		}

		if matrix[row][line] == target {
			return true
		}

		if target < matrix[row][line] && row > 0 && matrix[row-1][line] >= target {
			row--
			continue
		}
		line--
	}
	return false
}
