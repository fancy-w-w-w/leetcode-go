package doubleptr

import "math"

type TwoSum struct {
	val map[int]int
	max int
	min int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {

	val := make(map[int]int)
	return TwoSum{val: val, min: math.MaxInt64, max: math.MinInt64}

}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {

	this.min = min(this.min, number)
	this.max = max(this.max, number)
	this.val[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {

	if value < this.min+this.min || value > this.max+this.max {
		return false
	}
	for k, v := range this.val {
		if value-k == k {
			if v > 1 {
				return true
			}

		} else {
			if this.val[value-k] > 0 {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
