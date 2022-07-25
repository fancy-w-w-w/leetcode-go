package main

import (
	"fmt"
	"time"

	"github.com/umpc/go-sortedmap"
	"github.com/umpc/go-sortedmap/asc"
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
	// Create an empty SortedMap with a size suggestion and a less than function:
	sm := sortedmap.New(4, asc.Time)

	// Insert example records:
	sm.Insert("OpenBSD", time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC))
	sm.Insert("UnixTime", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))
	sm.Insert("Linux", time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC))
	sm.Insert("GitHub", time.Date(2008, 4, 10, 0, 0, 0, 0, time.UTC))

	// Set iteration options:
	reversed := false
	lowerBound := time.Date(1994, 1, 1, 0, 0, 0, 0, time.UTC)
	upperBound := time.Now()

	// Select values > lowerBound and values <= upperBound.
	// Loop through the values, in reverse order:
	iterCh, err := sm.BoundedIterCh(reversed, lowerBound, upperBound)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer iterCh.Close()

	for rec := range iterCh.Records() {
		fmt.Printf("%+v\n", rec)
	}
}
