package arrayhandle

import "math"

// 丑数
// 把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第 n个丑数。
func GetUglyNumber_Solution(index int) int {
	// write code here
	if index <= 6 {
		return index
	}

	i2, i3, i5 := 0, 0, 0
	res := make([]int, index)
	// 第一个丑数
	res[0] = 1

	for i := 1; i < index; i++ {
		//         var index int
		res[i], _ = minNum(res[i2]*2, res[i3]*3, res[i5]*5)
		if res[i] == res[i2]*2 {
			i2++
		}
		if res[i] == res[i3]*3 {
			i3++
		}
		if res[i] == res[i5]*5 {
			i5++
		}
	}
	return res[index-1]
}

func minNum(nums ...int) (minNum, index int) {
	minNum = math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < minNum {
			minNum = nums[i]
			index = i
		}
	}
	return
}
