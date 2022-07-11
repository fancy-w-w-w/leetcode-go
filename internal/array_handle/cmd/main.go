package main

import (
	"fmt"
	"strings"
)

type person struct {
	name  []byte
	score int
}

func main() {
	// n, isUp := 0, 0
	// fmt.Scanln(&n)
	// fmt.Scanln(&isUp)
	// input := bufio.NewScanner(os.Stdin)
	// input.Split(bufio.ScanWords)
	// res := []*person{}
	// for i := 0; i < n; i++ {
	// 	input.Scan()
	// 	name := input.Bytes()
	// 	input.Scan()
	// 	score, _ := strconv.Atoi(string(input.Bytes()))
	// 	res = append(res, &person{name: name, score: score})
	// }
	// if isUp == 0 {
	// 	sort.SliceStable(res, func(i, j int) bool {
	// 		return res[i].score > res[j].score
	// 	})
	// } else {
	// 	sort.SliceStable(res, func(i, j int) bool {
	// 		return res[i].score < res[j].score
	// 	})
	// }

	// for i := range res {
	// 	fmt.Printf("%s %d\n", res[i].name, res[i].score)
	// }
	// n := [2][4]int{}
	// for j := 0; j < 2; j++ {
	// 	for i := 0; i < 4; i++ {
	// 		fmt.Scan(&n[j][i])
	// 	}
	// }

	// fmt.Println(n)

	// script := `
	// local scores = {}
	// while #ARGV > 0 do
	// 	scores[#scores+1] = redis.call('ZSCORE', KEYS[1], table.remove(ARGV, 1))
	// end
	// return scores`
	// fmt.Println(script)

	ss := &strings.Builder{}

	ss.WriteString("???")
	ss.WriteByte('w')
	fmt.Println(ss.String())
}
