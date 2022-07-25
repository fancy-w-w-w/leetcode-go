package arrayhandle

// dfs使用全局变量
func DecodeString(s string) string {
	i = 0
	return run(s)
}

var i int

func run(s string) string {
	res := ""
	num := 0
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] >= 'a' && s[i] <= 'z' {
			res += string(s[i])
		} else if s[i] == '[' {
			i++
			tmp := run(s)
			for j := 0; j < num; j++ {
				res += tmp
			}
			num = 0
		} else {
			break
		}
	}
	return res
}
