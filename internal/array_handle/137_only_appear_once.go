package arrayhandle

// 只出现一次的数字 II
// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
// 位运算记录结果位数是否出现过一次
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		// 如果该位上不为1，说明都是出现三次或者0次累加的
		if total%3 > 0 {
			// 如果该位上为1，证明只出现一次的元素该位上也为1
			ans |= 1 << i
		}
	}
	return int(ans)
}
