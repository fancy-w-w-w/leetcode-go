package dfs

import "sort"

// 划分为k个相等的子集
// 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等
// dfs + 剪枝
func CanPartitionKSubsets(nums []int, k int) bool {
	// 为了后面剪枝
	sort.Ints(nums)

	//递归法解题 复杂度 k*2^n
	n := len(nums)
	if k > n {
		return false
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%k != 0 {
		return false
	}

	target := sum / k
	used := make([]bool, n)

	var dfs func(int, int, int) bool

	dfs = func(start int, k int, cursum int) bool {

		if k == 0 {
			return true
		}

		if cursum > target {
			return false
		}

		if cursum == target {
			return dfs(0, k-1, 0)
		}

		for i := start; i < n; i++ {
			if used[i] {
				continue
			}

			//最精彩的一步，如果前一个数没用（不符合条件），如果当前数和前一个相等，那么这个数也没用，直接抛弃即可！
			if (i > 0) && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}

			if nums[i] > target {
				//used[i] = true
				continue

			}
			if cursum+nums[i] > target {
				continue
			}

			used[i] = true
			cursum += nums[i]

			if dfs(i+1, k, cursum) {
				return true
			}

			used[i] = false
			cursum -= nums[i]

		}
		return false
	}

	return dfs(0, k, 0)
}
