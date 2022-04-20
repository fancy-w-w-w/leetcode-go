package digit

// PlusOne 给数加一
func PlusOne(digits []int) []int {
	jinwei := 1
	res := digits
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		tmp := digits[i]
		digits[i] = (digits[i] + jinwei) % 10
		jinwei = (tmp + jinwei) / 10
	}

	if jinwei != 0 {
		newStore := []int{1}
		res = append(newStore, digits...)
	}
	return res
}
