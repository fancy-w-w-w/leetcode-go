package main

import (
	"fmt"
	"log"
	"os/exec"
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

// func main() {
// 	a, b, c := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(3)
// 	go f("a", a, b, wg)
// 	go f("b", b, c, wg)
// 	go f("c", c, a, wg)
// 	a <- struct{}{}
// 	wg.Wait()
// }

// func f(s string, in, out chan struct{}, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 5; i++ {
// 		_ = <-in
// 		println(s)
// 		out <- struct{}{}
// 	}
// }

func main() {
	cmd := exec.Command("sh", "-c", "/home/tigerwang/work/nio-soa-converter-tool/resource/fidl-tool/fidl-tools-generator-linux-x86_64 -sf fidl -tf proto -sp /home/tigerwang/work/nio-soa-converter-tool/data/sourcefile/FrontBeam.fidl -d ./ -dv")
	// cmd := exec.Command("sh", "-c", "ls", "-a")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
