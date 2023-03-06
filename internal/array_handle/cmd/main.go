package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nums := sc.Bytes()
	var k int
	fmt.Scan(&k)

	window := []byte{}
	res := 0
	left := 0
	for i := range nums {
		window = append(window, nums[i])
		var num int
		// var err error
		num, _ = strconv.Atoi(string(window))
		for ; left <= i && num >= k; left++ {
			if len(window) == 0 {
				continue
			}
			window = window[1:]
			num, _ = strconv.Atoi(string(window))
		}
		res += i - left + 1
	}
	fmt.Println(res)
}

// func main() {
// 	var n, q int
// 	fmt.Scan(&n, &q)
// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&nums[i])
// 	}

// 	for i := 0; i < q; i++ {
// 		var handle, index, x int
// 		fmt.Scan(&handle, &index, &x)
// 		index--
// 		if handle == 1 {
// 			nums[index] = x
// 			// fmt.Println("??", index, nums)
// 		} else {
// 			var res int
// 			for j := 0; j <= index; j++ {
// 				if nums[j] == x {
// 					res++
// 				}
// 			}
// 			fmt.Println(res)
// 		}
// 	}
// }

// type person struct {
// 	name  []byte
// 	score int
// }

// func main() {
// 	// n, isUp := 0, 0
// 	// fmt.Scanln(&n)
// 	// fmt.Scanln(&isUp)
// 	// input := bufio.NewScanner(os.Stdin)
// 	// input.Split(bufio.ScanWords)
// 	// res := []*person{}
// 	// for i := 0; i < n; i++ {
// 	// 	input.Scan()
// 	// 	name := input.Bytes()
// 	// 	input.Scan()
// 	// 	score, _ := strconv.Atoi(string(input.Bytes()))
// 	// 	res = append(res, &person{name: name, score: score})
// 	// }
// 	// if isUp == 0 {
// 	// 	sort.SliceStable(res, func(i, j int) bool {
// 	// 		return res[i].score > res[j].score
// 	// 	})
// 	// } else {
// 	// 	sort.SliceStable(res, func(i, j int) bool {
// 	// 		return res[i].score < res[j].score
// 	// 	})
// 	// }

// 	// for i := range res {
// 	// 	fmt.Printf("%s %d\n", res[i].name, res[i].score)
// 	// }
// 	// n := [2][4]int{}
// 	// for j := 0; j < 2; j++ {
// 	// 	for i := 0; i < 4; i++ {
// 	// 		fmt.Scan(&n[j][i])
// 	// 	}
// 	// }

// 	// fmt.Println(n)

// 	// script := `
// 	// local scores = {}
// 	// while #ARGV > 0 do
// 	// 	scores[#scores+1] = redis.call('ZSCORE', KEYS[1], table.remove(ARGV, 1))
// 	// end
// 	// return scores`
// 	// fmt.Println(script)
// 	fmt.Println(f1([]int{1, 3, 4, 3, 3, 2, 4, 4}, 4))
// }

// func f1(arr []int, v int) int {
// 	left := 0
// 	right := len(arr) - 1
// 	for left <= right {
// 		mid := (left + right) / 2
// 		if arr[mid] < v {
// 			left = mid + 1
// 		} else if arr[mid] > v {
// 			right = mid - 1
// 		} else {
// 			return mid
// 		}
// 	}
// 	return -1
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
