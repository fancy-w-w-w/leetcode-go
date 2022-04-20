package matrix

// SetZeros 矩阵置零 让所有的与0相连的行列都置零
func SetZeros(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	// 标记第一列有没有0，如果有需要特殊处理
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
