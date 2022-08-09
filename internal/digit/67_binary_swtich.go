package digit

import "strconv"

// AddBinary 二进制求和
func AddBinary(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	str := ""

	flag := 0

	for len(a) > len(b) {
		b = "0" + b
	}

	for i := len(a) - 1; i >= 0; i-- {
		aNum := a[i] - '0'
		bNum := b[i] - '0'

		sum := int(aNum+bNum) + flag
		if sum == 2 {
			str = "1" + str
			flag = 1
		} else if sum == 3 {
			str = "1" + str
			flag = 1
		} else {
			str = strconv.Itoa(sum) + str
			flag = 0
		}
	}
	//确保最后一位需要进位
	if flag == 1 {
		str = "1" + str
	}
	return str

}
