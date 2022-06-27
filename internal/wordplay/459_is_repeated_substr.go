package wordplay

func RepeatedSubstringPattern(s string) bool {
	nextArr := kmpNextStr(s)
	return nextArr[len(s)-1] > 0 && len(s)%(len(s)-nextArr[len(s)-1]) == 0
}

func kmpNextStr(t string) []int {
	n := len(t)
	nextArray := make([]int, n)
	nextArray[0] = 0
	j := 0
	for i := 1; i < n; i++ {
		for j > 0 && t[i] != t[j] {
			j = nextArray[j-1]
		}
		if t[i] == t[j] {
			j++
		}
		nextArray[i] = j
	}
	return nextArray
}
