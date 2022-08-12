package dfs

import "sort"

// Combine 全排列组合，【1，n】生成长度为k的排列数组
// 容易踩坑，[][]int不能简单append(深拷贝)
func Combine(n, k int) [][]int {
	res := make([][]int, 0)
	roundElem := make([]int, 0)
	var dfs func(nowIndex int)
	dfs = func(nowIndex int) {
		if len(roundElem) == k {
			res = append(res, append([]int{}, roundElem...))
			return
		}

		for i := nowIndex; i <= n; i++ {
			roundElem = append(roundElem, i)
			dfs(i + 1)
			roundElem = roundElem[:len(roundElem)-1]
		}

	}

	dfs(1)
	return res
}

// 全排列2
// 去重
func PermuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
