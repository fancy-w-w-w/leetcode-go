package arrayhandle

//
func RemoveDuplicates(s string) string {
	if len(s) == 0 {
		return s
	}
	stack := make([]byte, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		str := s[i]
		if len(stack) == 0 || stack[len(stack)-1] != str {
			stack = append(stack, str)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	var res []byte
	for i := range stack {
		res = append(res, stack[i])
	}
	return string(res)
}
