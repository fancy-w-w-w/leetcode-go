package arrayhandle

func CountSmaller(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = partition2(nums, i)
	}
	return res
}

func partition2(arr []int, pivot int) int {
	l, r := 0, len(arr)-1
	arr[l], arr[pivot] = arr[pivot], arr[l]

	v := arr[l]
	i, j := l+1, r
	for {
		// 左边找到第一个不小于v的元素
		for i <= r && arr[i] < v {
			i++
		}
		// 右边找到第一个不大于v的元素
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
