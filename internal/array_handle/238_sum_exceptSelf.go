package arrayhandle

// ProductExceptSelf 除自身以外数组的乘积 前缀和知识
func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	// 前缀乘积
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// 后缀乘积
	postSum := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postSum
		postSum *= nums[i]
	}

	return res
}
