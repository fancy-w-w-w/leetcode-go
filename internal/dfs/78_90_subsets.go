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

func SubSetsWithoutDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
outer:
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				// 被子集选中且值相同 跳过
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}
