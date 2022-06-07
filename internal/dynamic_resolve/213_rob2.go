package dynamicresolve

// RobV2 所有的房屋都 围成一圈 ,这意味着第一个房屋和最后一个房屋是紧挨着的.
func RobV2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(RobV1(nums[:len(nums)-1]), RobV1(nums[1:len(nums)]))
}
