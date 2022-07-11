package arrayhandle

import "sort"

// MinDeletions 字符频次唯一的最小删除次数
func MinDeletions(s string) int {
	cnt := make([]int, 26)
	for _, r := range s {
		cnt[int(r-'a')]++
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	res := 0
	for i := 1; i < 26; i++ {
		for cnt[i] > 0 && cnt[i] >= cnt[i-1] {
			res++
			cnt[i]--
		}
	}
	return res
}
