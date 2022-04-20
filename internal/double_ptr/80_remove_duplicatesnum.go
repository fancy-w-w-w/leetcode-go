package doubleptr

import "fmt"

// RemoveDuplicates 每个元素最多出现k次，有序数组
// return 数组长度
func RemoveDuplicates(nums []int, k int) int {
	n := len(nums)
	if n <= k {
		return n
	}
	slow, fast := k, k
	for fast < n {
		if nums[slow-k] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	fmt.Println(nums[:slow])
	return slow
}
