package ms

import "fmt"

// 3. 好子串的个数，好子串的定义为一个子串，
// 且其中只有一个字母出现的次数为1次，其余字母均出现偶数次，给定一个字符串，求好子串的个数
func MayiFunc3(s string) int64 {
	map1 := make(map[int]int, 0)
	var res int64 = 0
	cur := 0
	map1[0] = 1
	for i := 0; i < len(s); i++ {
		cur ^= 1 << int(s[i]-'a')
		for i := 0; i < 26; i++ {
			res += int64(map1[cur^(1<<i)])
		}
		tmp := map1[cur] + 1
		map1[cur] = tmp
		fmt.Println(cur)
	}

	return res
}
