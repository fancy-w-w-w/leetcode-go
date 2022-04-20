package dfs

// Combine 全排列组合，【1，n】生成长度为k的排列数组
// 容易踩坑，[][]int不能简单append(深拷贝)
func Combine(n, k int) [][]int {
	res := make([][]int, 0)
	roundElem := make([]int, 0)
	var dfs func(nowIndex int)
	dfs = func(nowIndex int) {
		if len(roundElem) == k {
			res = append(res, append([]int{}, roundElem...))
			return
		}

		for i := nowIndex; i <= n; i++ {
			roundElem = append(roundElem, i)
			dfs(i + 1)
			roundElem = roundElem[:len(roundElem)-1]
		}

	}

	dfs(1)
	return res
}
