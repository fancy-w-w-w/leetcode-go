package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串划分
func main1() {
	shuru := "toyxmy"
	numx, numy := 0, 0
	rnumx, rnumy := 0, 0
	for i := range shuru {
		if shuru[i] == 'x' {
			rnumx++
		} else if shuru[i] == 'y' {
			rnumy++
		}
	}
	res := 0
	for i := 0; i < len(shuru); i++ {
		if (numx == numy || rnumx == rnumy) && i != 0 {
			res++
		}
		if shuru[i] == 'x' {
			numx++
			rnumx--
		} else if shuru[i] == 'y' {
			numy++
			rnumy--
		}
	}
	fmt.Println(res)
}

// 网格移动问题 不会
func main2() {

}

// 求大于N并且没有连续数字的数
func main() {
	shuru := 44432
	var judgeFunc func(str string) (result string, isOver bool)
	judgeFunc = func(str string) (result string, isOver bool) {
		isOver = true
		for i := range str {
			if i == 0 {
				continue
			}
			n := len(str)
			if str[i] == str[i-1] {
				head := string(str[:i+1] + strings.Repeat("0", n))
				headInt, _ := strconv.Atoi(head)
				rearInt, _ := strconv.Atoi(str[i-1 : i+1])
				result = strconv.Itoa(headInt + rearInt)
				isOver = false
			}
		}
		if !isOver {
			result, isOver = judgeFunc(result)
		}
		return
	}

	fmt.Println(judgeFunc(strconv.Itoa(shuru)))
}
