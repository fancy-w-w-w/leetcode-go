package arrayhandle

func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

// 存在重复元素3
// 给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。

// 如果存在则返回 true，不存在返回 false。

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 我们也可以使用利用桶排序的思想解决本题。我们按照元素的大小进行分桶，维护一个滑动窗口内的元素对应的元素。

// 对于元素 xx，其影响的区间为 [x - t, x + t][x−t,x+t]。于是我们可以设定桶的大小为 t + 1t+1。如果两个元素同属一个桶，那么这两个元素必然符合条件。如果两个元素属于相邻桶，那么我们需要校验这两个元素是否差值不超过 tt。如果两个元素既不属于同一个桶，也不属于相邻桶，那么这两个元素必然不符合条件。

// 具体地，我们遍历该序列，假设当前遍历到元素 xx，那么我们首先检查 xx 所属于的桶是否已经存在元素，如果存在，那么我们就找到了一对符合条件的元素，否则我们继续检查两个相邻的桶内是否存在符合条件的元素。

// 实现方面，我们将 \texttt{int}int 范围内的每一个整数 xx 表示为 x = (t + 1) \times a + b(0 \leq b \leq t)x=(t+1)×a+b(0≤b≤t) 的形式，这样 xx 即归属于编号为 aa 的桶。因为一个桶内至多只会有一个元素，所以我们使用哈希表实现即可。
