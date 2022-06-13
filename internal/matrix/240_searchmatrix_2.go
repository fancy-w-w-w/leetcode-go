package matrix

// SearchMatrixV2 搜索全局有序的数组
// 从右上角开始
func SearchMatrixV2(matrix [][]int, target int) bool {
	row, line := 0, len(matrix[0])-1

	for row < len(matrix) && line >= 0 {
		val := matrix[row][line]
		if val == target {
			return true
		} else if val > target {
			line--
		} else {
			row++
		}
	}
	return false
}
