package dfs

import "sort"

// 组合总数1
// 可以重复选取
func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

// 组合总数2
// candidates 中的每个数字在每个组合中只能使用 一次
// 有重复元素
// ：解集不能包含重复的组合
func CombinationSum2(candidates []int, target int) [][]int {
	// sort
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	res := make([][]int, 0)
	oneTurnSlice := make([]int, 0)
	var dfs func(target int, index int)
	dfs = func(target int, index int) {
		if target == 0 {
			res = append(res, append([]int{}, oneTurnSlice...))
			return
		}
		for i := index; i < len(candidates); i++ {
			// sort过后直接返回
			if target-candidates[i] < 0 || (i > index && candidates[i] == candidates[i-1]) {
				continue
			}
			oneTurnSlice = append(oneTurnSlice, candidates[i])
			dfs(target-candidates[i], i+1)
			oneTurnSlice = oneTurnSlice[:len(oneTurnSlice)-1]
		}
	}
	dfs(target, 0)
	return res
}
