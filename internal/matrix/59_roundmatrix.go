package matrix

func GenerateMatrix(n int) [][]int {
	res := make([][]int, n)

	for index, _ := range res {
		res[index] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1

	currentNum := 1
	target := n * n

	for currentNum <= target {
		for i := left; i <= right; i++ {
			res[top][i] = currentNum
			currentNum++
		}
		top++

		for j := top; j <= bottom; j++ {
			res[j][right] = currentNum
			currentNum++
		}
		right--

		for k := right; k >= left; k-- {
			res[bottom][k] = currentNum
			currentNum++
		}
		bottom--

		for v := bottom; v >= top; v-- {
			res[v][left] = currentNum
			currentNum++
		}
		left++
	}

	return res
}
