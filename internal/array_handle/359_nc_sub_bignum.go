package arrayhandle

// 两数相减
// 字符串
func SubTwostring(num1 string, num2 string) string {
	// write code here
	//条件一：num1 && num2 不可能等于0
	//条件二：两个数字都为正整数
	//其实就是，简单的大数减小数，小数减去大数

	len1 := len(num1)
	len2 := len(num2) //假设为最大值

	// 信号器：若num1 大于 num2 则是 大数 减去 小数 为 正整数
	var flag bool = false
	if len1 > len2 {
		len1, len2 = len2, len1
		num1, num2 = num2, num1
		flag = true
	}

	sum := make([]byte, len2) //存储计算的值(减法有可能负负得正)
	var cut int = 0           //借位
	for i := 0; i < len2; i++ {
		v2 := int(num2[len2-i-1] - '0') //num2的倒数第一个值
		// 没有的话为零
		var v1 int = 0
		if len1-i-1 >= 0 {
			v1 = int(num1[len1-i-1] - '0') //num1的倒数第一个值
		}

		result := v2 - v1 - cut
		if result < 0 {
			cut = 1
			sum[len2-i-1] = byte(result + 10 + '0')
			continue
		}
		cut = 0
		sum[len2-i-1] = byte(result + '0')

		// 3278729808879
		//  429018691214
	}
	// 因为100 -1 = 99 ；101 - 1 = 100 ；这个位数是变动的，以此来统计0的数量，之间截取就可以
	zeroCount := 0
	for j := 0; j < len(sum); j++ {
		// 连续的0，则为计算的空值0
		if sum[j] == '0' {
			zeroCount += 1
			continue
		}
		// 若不是连续的0，说明是正常值
		break
	}
	// 若0的数量与存储长度一致，说明结果为0
	if zeroCount == len(sum) {
		return "0"
	}

	// 大数 - 小数 = 正数
	if flag {
		return string(sum[zeroCount:])
	}

	// 小数 - 大数 = 负数
	return "-" + string(sum[zeroCount:])

}
