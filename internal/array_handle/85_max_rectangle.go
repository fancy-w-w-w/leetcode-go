package arrayhandle

// func maximalRectangle(matrix [][]byte) (ans int) {
// 	if len(matrix) == 0 {
// 		return
// 	}
// 	m, n := len(matrix), len(matrix[0])
// 	left := make([][]int, m)
// 	for i, row := range matrix {
// 		left[i] = make([]int, n)
// 		for j, v := range row {
// 			if v == '0' {
// 				continue
// 			}
// 			if j == 0 {
// 				left[i][j] = 1
// 			} else {
// 				left[i][j] = left[i][j-1] + 1
// 			}
// 		}
// 	}
// 	for i, row := range matrix {
// 		for j, v := range row {
// 			if v == '0' {
// 				continue
// 			}
// 			width := left[i][j]
// 			area := width
// 			for k := i - 1; k >= 0; k-- {
// 				width = min(width, left[k][j])
// 				area = max(area, (i-k+1)*width)
// 			}
// 			ans = max(ans, area)
// 		}
// 	}
// 	return
// }

// 最大矩形
// 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
// 单调栈
// 先把每一列的1的数量统计出来
func MaximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	heights := make([]int, n)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 只要出现 '0'，高度就要置0，否则在原来的基础上+1
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		ans = max(ans, LargestRectangleArea(heights))
	}
	return ans
}
