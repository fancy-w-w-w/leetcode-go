package doubleptr

// SortColors 相同颜色（0,1,2）的元素从小到大排序，原地算法
func SortColors(nums []int) []int {
	n := len(nums)

	if n < 2 {
		return nums
	}

	// 双指针，0\2元素双指针，1元素居中不动
	ptr1, ptr2 := 0, n-1
	// 上限为ptr2
	for i := 0; i <= ptr2; i++ {
		if nums[i] == 0 {
			nums[i] = nums[ptr1]
			nums[ptr1] = 0
			ptr1++
		}

		if nums[i] == 2 {
			nums[i] = nums[ptr2]
			nums[ptr2] = 2
			ptr2--
			// 重点，回退一个指针，重新判断互换过后该位置的元素
			if nums[i] != 1 {
				i--
			}
		}

	}
	return nums
}
