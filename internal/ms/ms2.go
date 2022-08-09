package ms

// 组合数
// func SearchCombineNum(nums []int, target int) string {
// 	sort.Ints(nums)
// 	// flag
// 	flag := false
// 	targetStr := strconv.Itoa(target)
// 	res := []byte{}
// 	for i := 0; i < len(targetStr); i++ {
// 		targetChr := targetStr[i] - '0'
// 		fmt.Println(targetChr, flag)
// 		if !flag {
// 			l := binarySearch(nums, int(targetChr))
// 			if nums[l] > target {
// 				res = []byte{}
// 				for j := len(targetStr) - 1; j > 0; j-- {
// 					res = append(res, byte(nums[len(nums)-1]+'0'))
// 				}
// 				return string(res)
// 			}
// 			if nums[l] < int(targetChr) {
// 				flag = true
// 			}
// 			res = append(res, byte('0'+nums[l]))
// 		} else {
// 			res = append(res, byte('0'+nums[len(nums)-1]))
// 		}
// 	}
// 	return string(res)
// }

// func binarySearch(nums []int, target int) int {
// 	l, r := 0, len(nums)-1
// 	for l < r {
// 		mid := l + r + 1>>1
// 		if nums[mid] <= target {
// 			l = mid
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return l
// }

// 组合数
// 小于n最大数
func SearchCombineNum(nums []int, target int) int {
	res := -1

	tmp, cnt := target, 0
	// 统计target的位数
	for tmp != 0 {
		cnt++
		tmp /= 10
	}

	// dfs函数
	var dfs func(lenCombine, curlen int, curNum int)
	dfs = func(lenCombine, curlen, curNum int) {
		if curlen == lenCombine {
			if curNum < target && curNum > res {
				res = curNum
			}
			return
		}

		for i := 0; i < len(nums); i++ {
			curNum = curNum*10 + nums[i]
			dfs(lenCombine, curlen+1, curNum)
			curNum /= 10
		}
	}

	// t是组合数的长度curlen，合法解的位数一定与target相同或者小于1，并且位数>=1
	for t := max(1, cnt-1); t <= cnt; t++ {
		dfs(t, 0, 0)
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
