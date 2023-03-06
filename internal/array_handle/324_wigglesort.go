package arrayhandle

// 摆动排序 II
// 给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
// 交替 桶排序
// 先从最大数填充摆动数组奇数元素，之后填充较小偶数元素
func wiggleSort(nums []int) {
	// 0 <= nums[i] <= 5000
	var bucket [5001]int
	for i := range nums {
		bucket[nums[i]]++
	}

	var index1 int = 5000
	// 回填
	for i := 1; i < len(nums); i += 2 {
		for bucket[index1] == 0 {
			index1--
		}
		nums[i] = index1
		bucket[index1]--
	}

	for j := 0; j < len(nums); j += 2 {
		for bucket[index1] == 0 {
			index1--
		}
		nums[j] = index1
		bucket[index1]--
	}
}
