package wordplay

import "strings"

func LengthOfLastWord(s string) int {
	ss := strings.Split(s, " ")

	return len(ss[len(ss)-1])
}
