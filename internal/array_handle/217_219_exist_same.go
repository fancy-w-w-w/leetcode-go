package arrayhandle

func ContainsNearbyDuplicate(nums []int, k int) bool {
	set := map[int]struct{}{}
	for i, num := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
