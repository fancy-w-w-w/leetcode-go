package arrayhandle

// 找到所有数组中消失的数字
// 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
func FindDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		// 减一是填充数组下标从零开始
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}
