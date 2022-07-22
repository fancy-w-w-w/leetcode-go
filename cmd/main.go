package main

import (
	"fmt"
	"math"
	"math/rand"
)

// func main() {
// 	var x, y, a, b int
// 	fmt.Scanf("%d%d%d%d", &x, &y, &a, &b)

// 	if b == 0 || b%a != 0 {
// 		fmt.Println(-1)
// 	}
// 	fmt.Println(x, y, a, b)
// 	res := math.MaxInt
// 	var dfs func(int, int)
// 	dfs = func(j int, multiCount int) {
// 		if j > b {
// 			return
// 		}
// 		if j == b {
// 			res = min(multiCount, res)
// 			return
// 		}
// 		fmt.Println(j)
// 		dfs(j*x, multiCount+1)
// 		dfs(j*y, multiCount+1)
// 	}

// 	dfs(a, 0)
// 	if res == math.MaxInt {
// 		fmt.Println(-1)
// 		return
// 	}
// 	fmt.Println(res)

// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	var n, m int
// 	fmt.Scanf("%d%d", &n, &m)
// 	matrix := make([][]int, n)
// 	for j := 0; j < n; j++ {
// 		matrix[j] = make([]int, m)
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			fmt.Scanf("%d", &matrix[i][j])
// 		}
// 	}

// 	res := 0
// 	round := min(n, m)
// 	for k := 2; k <= round; k++ {
// 		for i := 1; i < n; i++ {
// 			for j := 1; j < m; j++ {
// 				if i-k+1 < 0 || j-k+1 < 0 {
// 					continue
// 				}
// 				tmp := matrix[i-k+1][j-k+1] + matrix[i-k+1][j] + matrix[i][j-k+1] + matrix[i][j]
// 				res = max(res, tmp)

// 			}
// 		}
// 	}

// 	fmt.Println(res)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int63n(int64(math.Pow10(10))*9) + int64(math.Pow10(10)))
	}
}
