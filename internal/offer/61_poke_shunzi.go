package offer

import "sort"

// 判断扑克牌顺子
// 从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

func IsStraight(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++ //记录joker数量
		} else {
			if nums[i] == nums[i+1] { //除了0之外，有相同的元素直接返回false
				return false
			}
		}
	}
	return nums[4]-nums[joker] < 5
}
