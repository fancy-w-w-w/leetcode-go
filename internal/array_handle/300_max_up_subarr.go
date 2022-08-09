package arrayhandle

//贪心+二分查找，根据官方题解的Golang版
func lengthOfLIS(nums []int) int {
	max := []int{nums[0]} //贪心数组，用于记录长度为i的子序列末尾的最小元素值
	var l, r, mid int
	for i := 1; i < len(nums); i++ {
		length := len(max)
		l, r = 0, length-1
		if nums[i] > max[length-1] {
			max = append(max, nums[i])
		} else if nums[i] < max[length-1] {
			//二分查找
			for l < r {
				mid = (l + r) / 2
				if max[mid] == nums[i] { //遇到max数组已有的值需跳过
					break
				} else if nums[i] < max[mid] {
					r = mid
				} else {
					//小技巧：由于mid向下取整，l取mid+1，最后l就是需要改变的max数组下标
					l = mid + 1
				}
			}
			//遇到max数组已有的值需跳过
			if max[mid] == nums[i] {
				continue
			}
			max[l] = nums[i]
		}
	}
	return len(max)
}
