package main

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
func main() {

}
