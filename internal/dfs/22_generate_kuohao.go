package dfs

// 括号生成
func GenerateParenthesis(n int) []string {
	res := []string{}
	var dfs func(path []byte, leftNum, rightNum int)
	dfs = func(path []byte, leftNum, rightNum int) {
		if len(path) == 2*n {
			res = append(res, string(path))
		}
		if rightNum > leftNum {
			return
		}

		if leftNum < n {
			path = append(path, '(')
			dfs(path, leftNum+1, rightNum)
			path = path[:len(path)-1]
		}

		if rightNum < leftNum {
			path = append(path, ')')
			dfs(path, leftNum, rightNum+1)
			path = path[:len(path)-1]
		}
	}

	dfs([]byte{}, 0, 0)
	return res
}
