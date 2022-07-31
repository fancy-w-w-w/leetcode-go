package offer

// 剑指 Offer 04. 二维数组中的查找
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix)
	m := len(matrix[0])
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}

	return false
}
