package arrayhandle

// 分割等和子集
// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
func canPartition(nums []int) bool {
	/**
	  动态五部曲：
	      1.确定dp数组和下标含义
	      2.确定递推公式
	      3.dp数组初始化
	      4.dp遍历顺序
	      5.打印
	  **/
	//确定和
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 { //如果和为奇数，则不可能分成两个相等的数组
		return false
	}
	sum /= 2
	//确定dp数组和下标含义
	var dp [][]bool //dp[i][j] 表示： 前i个石头是否总和不大于J
	//初始化数组
	dp = make([][]bool, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= sum; j++ { //j是固定总量
			if j >= nums[i-1] { //如果容量够用则可放入背包
				// 一种是将nums[i-1]添加进来，另一种是不添加
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else { //如果容量不够用则不拿，维持前一个状态
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][sum]
}
