package dfs

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
