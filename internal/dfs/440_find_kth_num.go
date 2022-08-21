package dfs

//  字典序第K小数字
// 给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。
// 字典序[1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]
func FindKthNumber(n int, k int) int {
	// 前缀为 prefix 时，子树还有多少元素，prefix 用来表示每一层的起始数字
	countNodes := func(prefix int) int {
		width := 1           // 当前层的宽度
		cnt := 0             // 节点总数
		for prefix*10 <= n { // 如果下一层起始节点值小于等于 n，本层节点可以全部算入
			cnt += width // 节点总数为每一层节点数之和
			width *= 10  // 每一层节点数翻 10 倍
			prefix *= 10 // 更新起始节点
		}
		cnt += min(width, n-prefix+1) // 加上最后一层的节点数
		return cnt
	}

	begin := 1 // 所有情况的起始节点都为 1

	// 一直迭代到 k == 1，因为某个节点下的第一个最小值即为节点本身
	for k > 1 {
		t := countNodes(begin) // 以 begin 为根的节点数
		if t < k {
			// 说明第 k 小数不在以 begin 为根的节点下
			k -= t
			begin++
		} else {
			// 说明第 k 小数在以 begin 为根的节点下
			begin *= 10
			k--
		}
	}
	return begin
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
