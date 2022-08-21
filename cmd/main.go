package main

import (
	"fmt"
	"sync"
)

func main() {
	// channel 初始化
	c := make(chan int, 10)
	// 用来 recevivers 同步事件的
	wg := sync.WaitGroup{}

	// sender（写端）
	go func() {
		// 入队
		c <- 1
		// ...
		// 满足某些情况，则 close channel
		close(c)
	}()

	// receivers （读端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ... 处理 channel 里的数据
			for v := range c {
				mm := v
				fmt.Println(mm)
			}
		}()
	}
	// 等待所有的 receivers 完成；
	wg.Wait()
}

// func main() {
// 	ss := Solution("0000")
// 	fmt.Println(ss)
// }
// func Solution(S string) string {
// 	// write your code in Go (Go 1.4)
// 	map1 := make(map[int]int, 10)
// 	for i := range S {
// 		map1[int(S[i]-'0')]++
// 	}
// 	res := []byte{}
// 	for i := 9; i >= 0; i-- {
// 		if v, ok := map1[i]; ok && v > 1 {
// 			num := v / 2
// 			v %= 2
// 			map1[i] = v
// 			ss := []byte(strings.Repeat(string([]byte{byte('0' + i)}), num))
// 			res = append(ss, res...)

// 		}
// 	}

// 	var isInsert bool = false
// 	for j := 9; j >= 0; j-- {
// 		if v, ok := map1[j]; ok && v == 1 {
// 			res = append(res, byte('0'+j))
// 			isInsert = true
// 			break
// 		}
// 	}

// 	if isInsert {
// 		res = append(res, res[:len(res)-1]...)
// 	} else {
// 		res = append(res, res...)
// 	}
// 	lenOrigin := len(res)

// 	ss1 := strings.TrimLeft(string(res), "0")
// 	ss2 := strings.TrimRight(ss1, "0")
// 	// 特殊 "0000"
// 	if len(ss2) == 0 && lenOrigin != 0 {
// 		ss2 = "0"
// 	}
// 	return ss2
// }

// func main() {
// 	X := []int{4, 8, 2, 2, 1, 4}
// 	Y := []int{1, 2, 3, 1, 2, 3}
// 	W := 3
// 	fmt.Println(Solution(X, Y, W))
// }

// func Solution(X []int, Y []int, W int) int {
// 	// write your code in Go (Go 1.4)
// 	if len(X) == 0 {
// 		return 0
// 	}
// 	sort.Ints(X)

// 	res := 0
// 	i := 0
// 	for i < len(X) {
// 		nextIndex := X[i] + W
// 		nexti := i
// 		j := i
// 		for ; j < len(X); j++ {
// 			if X[j] <= nextIndex {
// 				continue
// 			}
// 			break
// 		}
// 		nexti = j
// 		i = nexti
// 		res++
// 	}
// 	return res

// }

// 字符串输入
// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	data, _ := reader.ReadString(' ') // data永远带个最后休止符
// 	str := string(data)
// 	fmt.Println(str)
// 	return
// }

// 字符串输入
// func main() {
// 	scan := bufio.NewScanner(os.Stdin)
// 	scan.Scan()
// 	str := scan.Text()
// 	scan.Scan()
// 	str2 := scan.Text()
// 	fmt.Println(str)
// 	fmt.Println(str2)
// }

// func main() {

// }

// // 数字输入
// func main() {
// 	var N, m int
// 	fmt.Scanln(&N, &m)
// 	fmt.Println(N, m)
// }

// func main() {
// 	var N int
// 	nums := make([]int, 0)
// 	fmt.Scan(&N)
// 	for i := 0; i < N; i++ {
// 		var tmp int
// 		fmt.Scan(&tmp)
// 		nums = append(nums, tmp)
// 	}
// 	fmt.Println(nums)
// }

// func main() {
// 	N := solution("RW")
// 	fmt.Println(N)
// }

// func solution(S string) int {
// 	numR := 0
// 	for i := range S {
// 		if S[i] == 'R' {
// 			numR++
// 		}
// 	}

// 	min, mid := 0, numR/2
// 	sum, flag := 0, 0
// 	listin := []int{}

// 	for right := 0; right < len(S); right++ {
// 		if S[right] == 'R' {
// 			listin = append(listin, right)
// 		}

// 		if len(listin) == numR && flag == 0 {
// 			flag = 1
// 			for i := 0; i < len(listin); i++ {
// 				if i < mid {
// 					sum += listin[mid] - listin[i] - (mid - i)
// 				} else {
// 					sum += listin[i] - listin[mid] - (i - mid)
// 				}
// 			}
// 			min = sum
// 		}

// 		if len(listin) > numR && mid > 0 {
// 			if numR%2 == 0 {
// 				sum += -listin[mid] + listin[0] + listin[len(listin)-1] - listin[mid]
// 			} else {
// 				sum += -listin[mid] + listin[0] + listin[len(listin)-1] - listin[mid+1]
// 			}
// 			listin = listin[1:]
// 			if sum < min {
// 				min = sum
// 			}
// 		}
// 	}

// 	if min > int(math.Pow10(9)) {
// 		return -1
// 	}
// 	return min
// }

// func main() {
// 	N := Solution("...xxx..x....xxx.", 7)
// 	fmt.Println(N)
// }

// func Solution(S string, B int) int {
// 	res := 0

// 	queue := build(S)
// 	for queue.Len() != 0 && B > 1 {
// 		poll := queue.Pop().(int)
// 		num := min(B-1, poll)
// 		res += num
// 		B -= num + 1
// 		if num == poll {
// 			continue
// 		}
// 		queue.Push(poll - num)
// 	}
// 	return res
// }

// func build(S string) *Heap {
// 	heap1 := &Heap{array: []int{}}
// 	heap.Init(heap1)
// 	N := 0
// 	for l, r := 0, 0; r < len(S); r++ {
// 		if S[r] == 'x' {
// 			N++
// 			continue
// 		}
// 		if N != 0 {
// 			heap1.Push(N)
// 		}
// 		for l < r {
// 			if S[l] == 'x' {
// 				N--
// 			}
// 			l++
// 		}

// 	}
// 	if N != 0 {
// 		heap1.Push(N)
// 	}
// 	return heap1
// }

// // heap大根堆
// type Heap struct {
// 	array []int
// }

// func (h *Heap) Len() int {
// 	return len(h.array)
// }

// func (h *Heap) Less(i, j int) bool {
// 	return h.array[i] > h.array[j]
// }

// func (h *Heap) Swap(i, j int) {
// 	h.array[i], h.array[j] = h.array[j], h.array[i]
// }

// func (h *Heap) Push(x interface{}) {
// 	h.array = append(h.array, x.(int))
// }

// func (h *Heap) Pop() interface{} {
// 	res := h.array[len(h.array)-1]
// 	h.array = h.array[:len(h.array)-1]
// 	return res
// }

// func min(A, B int) int {
// 	if A < B {
// 		return A
// 	}
// 	return B
// }

// func main() {
// 	A := []int{2, 5, 6, 5}
// 	B := []int{5, 4, 2, 2}
// 	N := len(A)
// 	nn := solution(A, B, N, 8)
// 	fmt.Println(nn)
// }
// func solution(A, B []int, N, S int) int {
// 	if N > S {
// 		return 0
// 	}
// 	pos := make([]int, S+1)
// 	time := make([][]int, 0)
// 	for i := 0; i < N; i++ {
// 		time = append(time, []int{min(A[i], B[i]), max(A[i], B[i])})
// 	}
// 	sort.Slice(time, func(i, j int) bool {
// 		if time[i][0] == time[j][0] {
// 			return time[i][1] < time[j][1]
// 		}
// 		return time[i][0] <= time[j][0]
// 	})
// 	for i := 0; i < N; i++ {
// 		l, r := time[i][0], time[i][1]
// 		if pos[l] == 0 {
// 			pos[l] = 1
// 			continue
// 		}
// 		if pos[r] == 0 {
// 			pos[r] = 1
// 			continue
// 		}
// 		return 0
// 	}
// 	return 1
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// func max(A, B int) int {
// 	if A > B {
// 		return A
// 	}
// 	return B
// }

// func Slice(x any, less func(i, j int) bool) {
// 	rv := reflectValueOf(x)
// 	swap := reflectSwapper(x)
// 	length := rv.Len()
// 	quickSort_func(lessSwap{less, swap}, 0, length, maxDepth(length))
// }

// func Solution(A []int, B []int, S int) bool {
// 	n := len(A)
// 	if n > S {
// 		return false
// 	}
// 	shuru := [100001]int{}
// 	var dfs func(index int) bool
// 	dfs = func(index int) bool {
// 		if index >= n {
// 			return true
// 		}
// 		if shuru[A[index]] > 0 && shuru[B[index]] > 0 {
// 			return false
// 		}
// 		var judge1, judge2 bool
// 		for i := index; i < n; i++ {
// 			if shuru[A[index]] == 0 {
// 				shuru[A[index]]++
// 				judge1 = dfs(index + 1)
// 				shuru[A[index]]--
// 			}
// 			if shuru[B[index]] == 0 {
// 				shuru[B[index]]++
// 				judge2 = dfs(index + 1)
// 				shuru[B[index]]--
// 			}
// 		}
// 		return judge1 || judge2

// 	}
// 	// shuru := [100001]int{}
// 	res := dfs(0)
// 	fmt.Println(shuru)
// 	return res
// }

// func main() {
// 	nums := []int{3, 0, 5}
// 	fmt.Println(Solution(nums))
// }

// func Solution(A []int) int {
// 	// write your code in Go (Go 1.4)
// 	var totalWuran float32 = 0
// 	res := 0
// 	nums := make([]float32, len(A))
// 	for i := range A {
// 		nums[i] = float32(A[i])
// 		totalWuran += float32(A[i])
// 	}

// 	hp1 := &hp{
// 		array: nums,
// 	}
// 	heap.Init(hp1)

// 	nowWuran := totalWuran
// 	// 除2减半，作为跳出条件
// 	totalWuran /= 2
// 	for nowWuran > totalWuran {
// 		v := heap.Pop(hp1).(float32)
// 		nowWuran -= v / 2
// 		heap.Push(hp1, v/2)
// 		res++
// 	}

// 	return res
// }

// type hp struct {
// 	array []float32
// }

// func (h *hp) Len() int {
// 	return len(h.array)
// }

// func (h *hp) Less(i, j int) bool {
// 	return h.array[i] > h.array[j]
// }

// func (h *hp) Swap(i, j int) {
// 	h.array[i], h.array[j] = h.array[j], h.array[i]
// }

// func (h *hp) Push(x any) {
// 	h.array = append(h.array, x.(float32))
// }

// func (h *hp) Pop() any {
// 	res := h.array[len(h.array)-1]
// 	h.array = h.array[:len(h.array)-1]
// 	return res
// }

// func main() {

// }

// func Solution(X []int, Y []int) int {
// 	// write your code in Go (Go 1.4)

// 	dp := make([][][]float32, len(X)+1)
// 	for i := range dp {
// 		dp[i] = make([][]float32, len(Y)+1)
// 		for j := range dp[j] {
// 			dp[i][j] = make([]float32, len(X)+1)
// 		}
// 	}

// 	for j := range Y {
// 		dp[0][j][0] = float32(X[0] / Y[j])
// 	}

// 	for i := 1; i < len(X); i++ {
// 		for k := 0; k < i; k++ {
// 			for j := 0; j < len(Y); j++ {
// 				dp[k][j][i] = 1
// 			}
// 		}
// 	}
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	// x := []int{1, 2, 3, 1, 2, 12, 8, 4}
// 	// y := []int{5, 10, 15, 2, 4, 15, 10, 5}

// 	x := []int{1, 1, 1}
// 	y := []int{2, 2, 2}
// 	fmt.Println(Solution(x, y))
// }

// func Solution(X []int, Y []int) int {
// 	if len(X) == 0 {
// 		return 0
// 	}
// 	// 排序去冲
// 	// sort.Ints(X)
// 	sort.Ints(Y)
// 	visited := make([]bool, len(Y))
// 	dp := make([][]float32, len(X)+1)
// 	for i := range dp {
// 		dp[i] = make([]float32, len(Y)+1)
// 	}

// 	for i := range X {
// 		for j := range Y {
// 			dp[i][j] = float32(X[i]) / float32(Y[j])
// 		}
// 	}

// 	res := 0
// 	var dfs func(fenziIndex int, target float32)
// 	dfs = func(fenziIndex int, target float32) {
// 		if target == 0 {
// 			// fmt.Println(fenziIndex)
// 			res++
// 			return
// 		} else if target < 0 {
// 			fmt.Println(fenziIndex, target)
// 			return
// 		}
// 		if fenziIndex >= len(X) {
// 			return
// 		}

// 		for i := fenziIndex; i < len(X); i++ {
// 			if
// 			for j := range Y {
// 				if j > 0 && Y[j] == Y[j-1] {
// 					continue
// 				}
// 				if visited[j] {
// 					continue
// 				}
// 				visited[j] = true
// 				target -= dp[i][j]
// 				dfs(fenziIndex+1, target)
// 				visited[j] = false
// 				target += dp[i][j]
// 			}
// 			go dfs(fenziIndex+1, target)
// 		}
// 	}
// 	dfs(0, 1)
// 	return res
// }

// func main() {
// 	go func() {
// 		log.Println(http.ListenAndServe("localhost:6060", nil))
// 	}()

// 	time.Sleep(time.Second)
// 	c := make(chan bool)
// 	<-c
// }

// func main() {
// 	n := 60000000
// 	tt := time.Date(2008, time.January, 1, 0, 0, 0, 0, time.UTC)
// 	tt = tt.Add(time.Second * time.Duration(n))
// 	fmt.Println(tt.Format(time.UnixDate))
// }

// func main() {
// 	// lis1 := &linkedlist.ListNode{
// 	// 	Val: 1,
// 	// 	Next: &linkedlist.ListNode{
// 	// 		Val: 2,
// 	// 		Next: &linkedlist.ListNode{
// 	// 			Val:  3,
// 	// 			Next: nil,
// 	// 		},
// 	// 	},
// 	// }

// 	// lis2 := &linkedlist.ListNode{
// 	// 	Val: 2,
// 	// 	Next: &linkedlist.ListNode{
// 	// 		Val: 1,
// 	// 		Next: &linkedlist.ListNode{
// 	// 			Val:  3,
// 	// 			Next: nil,
// 	// 		},
// 	// 	},
// 	// }
// 	// res := specialquestion.SubtractionByLinkedlist(lis1, lis2)
// 	// for res != nil {
// 	// 	fmt.Println(res.Val)
// 	// 	res = res.Next
// 	// }
// 	http.HandleFunc("/", Test1)
// 	http.ListenAndServe(":4399", nil)
// }

// func Test1(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("success")
// }
