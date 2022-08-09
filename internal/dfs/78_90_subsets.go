package dfs

import "sort"

// SubSets求数组的子集
func SubSets(nums []int) [][]int {
	res := make([][]int, 0)

	singleRound := make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(nums) {
			res = append(res, append([]int{}, singleRound...))
			return
		}
		singleRound = append(singleRound, nums[start])
		dfs(start + 1)
		singleRound = singleRound[:len(singleRound)-1]
		dfs(start + 1)
	}
	dfs(0)
	return res
}

// SubSetsV2 求子集迭代法
//用1和0表示选择或者不选择该位上的元素，例如（1110）表示该子集为选择前三个元素，最后一个元素不选择。
func SubSetsV2(nums []int) (ans [][]int) {
	// 需要排序
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

// 子集不重复
// 包含重复元素 先排序
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	t := []int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1]
	}
	dfs(false, 0)
	return
}
