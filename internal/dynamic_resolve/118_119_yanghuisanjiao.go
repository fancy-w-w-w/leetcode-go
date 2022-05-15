package dynamicresolve

func GenerateYangHui(numsRow int) [][]int {
	ans := make([][]int, numsRow)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans

}

// GenerateYangHuiV2 获得指定层的杨辉三角数组
func GenerateYangHuiV2(numsRow int) []int {
	res := make([]int, numsRow+1)
	res[0] = 1
	for i := 1; i <= numsRow; i++ {
		// res[i-1] * (n-m+1)/m
		res[i] = res[i-1] * (numsRow - i + 1) / i
	}
	return res
}
