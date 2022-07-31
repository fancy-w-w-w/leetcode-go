package offer

func MinArray(numbers []int) int {
	res := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			res = numbers[i]
			return res
		}
	}
	return res
}
