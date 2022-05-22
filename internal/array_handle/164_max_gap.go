package arrayhandle

// MaximumGap 求无序数组最大间距 基数排序
func MaximumGap(nums []int) int {
	res := 0
	RadixSort(nums)
	for i := range nums {
		if i == 0 {
			continue
		}
		res = max(res, nums[i]-nums[i-1])
	}
	return res
}

// radixSort 基数排序
func RadixSort(seq []int) {
	// figures:位数,都是大于等于0的数
	var max = 0
	for _, s := range seq {
		if s > max {
			max = s
		}
	}
	// 最大循环次数（数组中数字最大位数）
	var maxFigures = 1
	for max/10 > 0 {
		maxFigures++
		max /= 10
	}

	figuresSort(seq, maxFigures)
}

func figuresSort(seq []int, mfg int) {
	num := 1
	for i := 0; i < mfg; i++ {
		var bucket [10][]int
		var result []int
		// 放入对应桶
		for _, s := range seq {
			n := s / num % 10
			bucket[n] = append(bucket[n], s)
		}

		// 从小到大收集
		for i := 0; i < 10; i++ {
			result = append(result, bucket[i]...)
		}

		for i := range seq {
			seq[i] = result[i]
		}
		num *= 10
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
